package server

import (
	"context"
	"errors"

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

	for _, id := range s.node.NodeIDs() {
		if id == s.node.ID() {
			continue
		}
		v, err := s.ReadIntWithDefault(id)
		if err != nil {
			return err
		}
		total += v
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
