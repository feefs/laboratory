package server

import (
	"broadcast/server/rpc"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func (s *Server) InitHandler(msg maelstrom.Message) error {
	if s.node.ID() == "n0" {
		go s.batchPropagate(150 * time.Millisecond)
	}
	return nil
}

func (s *Server) batchPropagate(freq time.Duration) {
	tick := make(chan struct{})
	go func() {
		for {
			time.Sleep(freq)
			tick <- struct{}{}
		}
	}()

	for {
		select {
		case message := <-s.state.input:
			s.state.batch.messages = append(s.state.batch.messages, message)
		case <-tick:
			s.propagate()
			s.state.batch.messages = []int64{}
		}
	}
}

func (s *Server) propagate() {
	id, err := GeneratePropagateID()
	if err != nil {
		log.Printf("Unable to generate Propagation ID: %v", err)
		return
	}

	messages := make([]int64, len(s.state.batch.messages))
	copy(messages, s.state.batch.messages)

	propagateReq := &PropagateReqBody{
		MessageBody:   maelstrom.MessageBody{Type: "propagate"},
		PropagationID: id,
		Messages:      messages,
	}
	for _, nid := range s.node.NodeIDs() {
		if nid == "n0" {
			continue
		}
		go rpc.Retry(s.node, nid, propagateReq)
	}
}
