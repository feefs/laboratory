package server

import (
	"context"
	"errors"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ReadRespBody struct {
	maelstrom.MessageBody
	Value int `json:"value"`
}

func (s *server) ReadHandler(msg maelstrom.Message) error {
	if len(msg.Src) == 0 {
		return errors.New("empty caller type")
	}
	switch msg.Src[0] {
	case 'n':
		return s.handleReadNode(msg)
	case 'c':
		return s.handleReadClient(msg)
	}
	return errors.New("unknown caller type")
}

func (s *server) handleReadNode(msg maelstrom.Message) error {
	v, err := s.ReadIntWithDefault(s.node.ID())
	if err != nil {
		return err
	}

	respBody := &ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Value:       v,
	}

	return s.node.Reply(msg, respBody)
}

func (s *server) handleReadClient(msg maelstrom.Message) error {
	total, err := s.ReadIntWithDefault(s.node.ID())
	if err != nil {
		return err
	}

	type result struct {
		value int
		err   error
	}
	// resultsChan must be buffered. otherwise, there will be a deadlock
	// when the goroutines block on sending and never call wg.Done()
	resultsChan := make(chan result, len(s.node.NodeIDs())-1)

	var wg sync.WaitGroup
	for _, id := range s.node.NodeIDs() {
		if id == s.node.ID() {
			continue
		}
		wg.Add(1)
		go func() {
			v, err := s.ReadIntWithDefault(id)
			resultsChan <- result{value: v, err: err}
			wg.Done()
		}()
	}
	wg.Wait()
	// buffered channels must be closed before iterating over them.
	// otherwise, gathering results will loop forever.
	close(resultsChan)

	results := []result{}
	for result := range resultsChan {
		results = append(results, result)
	}

	for _, result := range results {
		if result.err != nil {
			return result.err
		}
		total += result.value
	}

	respBody := &ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Value:       total,
	}

	return s.node.Reply(msg, respBody)
}

func (s *server) ReadIntWithDefault(key string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	v, err := s.kv.ReadInt(ctx, key) // Returned value is sequentially consistent, no synchronization needed

	if maelstrom.ErrorCode(err) == maelstrom.KeyDoesNotExist {
		return 0, nil
	} else {
		return v, err
	}
}
