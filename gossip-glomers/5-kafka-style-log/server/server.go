package server

import (
	"errors"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var ErrNotImplemented = errors.New("not implemented")

type message struct {
	offset int
	value  int
}
type offsets map[string]([]message)
type committedOffsets map[string]int

type server struct {
	node               *maelstrom.Node
	offsets            offsets
	offsetsmu          sync.Mutex
	committedOffsets   committedOffsets
	committedOffsetsMu sync.Mutex
}

func NewServer(node *maelstrom.Node) *server {
	return &server{node: node, offsets: make(offsets), committedOffsets: make(committedOffsets)}
}
