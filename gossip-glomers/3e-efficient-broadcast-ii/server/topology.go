package server

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func (s *Server) TopologyHandler(msg maelstrom.Message) (err error) {
	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "topology_ok"})
}
