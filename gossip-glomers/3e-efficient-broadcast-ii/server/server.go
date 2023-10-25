package server

import (
	"broadcast/server/state"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Server struct {
	Node  *maelstrom.Node
	State *state.State
}

func NewServer(node *maelstrom.Node) *Server {
	return &Server{
		Node:  node,
		State: state.NewState(),
	}
}
