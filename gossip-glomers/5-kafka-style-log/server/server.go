package server

import (
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var ErrNotImplemented = errors.New("not implemented")

type offsets map[string]int

type server struct {
	node *maelstrom.Node
}

func NewServer(node *maelstrom.Node) *server {
	return &server{node}
}
