package server

import (
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var propagating = false

func (s *server) InitHandler(msg maelstrom.Message) error {
	if propagating {
		return nil
	}

	go func() {
		for {
			select {
			case msg := <-s.mp.messagesChan:
				s.mp.messages = append(s.mp.messages, msg)
			case <-s.mp.prepareReadMessagesChan:
				resp := make([]int, len(s.mp.messages))
				copy(resp, s.mp.messages)
				s.mp.readMessagesChan <- resp
			}
		}
	}()

	if s.node.ID() == "n0" {
		tick := make(chan struct{})
		go func() {
			for {
				tick <- struct{}{}
				// setting the propagation frequency to slightly above the
				// maelstrom test's network delay results in the lowest latency
				time.Sleep(110 * time.Millisecond)
			}
		}()

		go func() {
			for {
				select {
				case propagation := <-s.pp.propagationChan:
					s.pp.propagations = append(s.pp.propagations, propagation)
				case <-tick:
					propagationMessages := make(map[string]([]int), len(s.node.NodeIDs())-1)
					for _, propagation := range s.pp.propagations {
						for _, id := range s.node.NodeIDs() {
							// if propagation.srcID is a node, avoid propagating the message since the node already has it
							if id == "n0" || propagation.srcID == id {
								continue
							}
							propagationMessages[id] = append(propagationMessages[id], propagation.message)
						}
					}
					for _, id := range s.node.NodeIDs() {
						if id == "n0" {
							continue
						}
						go s.resilientRpc(id, &PropagateReqBody{
							MessageBody: maelstrom.MessageBody{Type: "propagate"},
							Messages:    propagationMessages[id],
						})
					}
					s.pp.propagations = []propagation{}
				}
			}
		}()
	}

	propagating = true

	return nil
}
