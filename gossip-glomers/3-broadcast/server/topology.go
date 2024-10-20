package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

func (s *server) TopologyHandler(msg maelstrom.Message) error {
	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "topology_ok"})
}
