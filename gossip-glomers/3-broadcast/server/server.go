package server

import (
	"context"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	node *maelstrom.Node
	mp   *messageProcessing
	pp   *propagationProcessing
}

type messageProcessing struct {
	messages                []int
	messagesChan            chan int
	prepareReadMessagesChan chan struct{}
	readMessagesChan        chan []int
}

type propagation struct {
	srcID   string
	message int
}

type propagationProcessing struct {
	propagationChan chan propagation
	propagations    []propagation
}

func (s *server) resilientRpc(id string, body any) {
	timeout := 1 * time.Second
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if _, err := s.node.SyncRPC(ctx, id, body); err == nil {
			break
		}
	}
}

func NewServer(node *maelstrom.Node) *server {
	return &server{
		node,
		&messageProcessing{
			messages:                []int{},
			messagesChan:            make(chan int),
			prepareReadMessagesChan: make(chan struct{}),
			readMessagesChan:        make(chan []int),
		},
		&propagationProcessing{
			propagationChan: make(chan propagation),
		},
	}
}
