package server

import (
	"broadcast/server/rpc"
	"broadcast/server/state"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func (s *Server) InitHandler(msg maelstrom.Message) error {
	if s.Node.ID() == "n0" {
		// setting the batch frequency to a little above the hardcoded network delay results in the lowest latency
		go s.batchPropagate(110 * time.Millisecond)
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
		case message := <-s.State.Batch.Input:
			s.State.Batch.Buffer = append(s.State.Batch.Buffer, message)
		case <-tick:
			// optimization: don't propagate if the batch buffer is empty
			if len(s.State.Batch.Buffer) > 0 {
				s.propagate()
				s.State.Batch.Buffer = []int64{}
			}
		}
	}
}

func (s *Server) propagate() {
	id, err := state.GeneratePropagateID()
	if err != nil {
		log.Printf("Unable to generate Propagation ID: %v", err)
		return
	}

	messages := make([]int64, len(s.State.Batch.Buffer))
	copy(messages, s.State.Batch.Buffer)

	propagateReq := &PropagateReqBody{
		MessageBody:   maelstrom.MessageBody{Type: "propagate"},
		PropagationID: id,
		Messages:      messages,
	}
	for _, nid := range s.Node.NodeIDs() {
		if nid == "n0" {
			continue
		}
		go rpc.Retry(s.Node, nid, propagateReq)
	}
}
