package server

import (
	"context"
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func (s *Server) InitHandler(msg maelstrom.Message) error {
	attempts := 0
	attempt_limit := 10
	for attempts < attempt_limit {
		ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		err := s.kv.Write(ctx, s.node.ID(), 0)
		if err == nil {
			break
		}
		attempts += 1
	}

	if attempts == attempt_limit {
		panic(errors.New("initial write failed"))
	}

	return nil
}
