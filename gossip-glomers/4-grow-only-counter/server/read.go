package server

import (
	"context"
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ReadRespBody struct {
	maelstrom.MessageBody
	Value int `json:"value"`
}

func (s *Server) ReadHandler(msg maelstrom.Message) error {
	if len(msg.Src) == 0 {
		return errors.New("empty caller type")
	}
	switch msg.Src[0] {
	case 'n':
		return s.handleReadNode(msg)
	case 'c':
		return s.handleReadClient(msg)
	}
	return errors.New("unknown caller type")
}

func (s *Server) handleReadNode(msg maelstrom.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	v, err := s.kv.ReadInt(ctx, s.node.ID()) // Returned value is sequentially consistent, no synchronization needed
	if err != nil {
		reqErr := &maelstrom.RPCError{}
		if errors.As(err, &reqErr) && reqErr.Code == maelstrom.KeyDoesNotExist {
			v = 0
		} else {
			return err
		}
	}

	respBody := &ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Value:       v,
	}

	return s.node.Reply(msg, respBody)
}

func (s *Server) handleReadClient(msg maelstrom.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	v, err := s.kv.ReadInt(ctx, s.node.ID()) // Returned value is sequentially consistent, no synchronization needed
	if err != nil {
		reqErr := &maelstrom.RPCError{}
		if errors.As(err, &reqErr) && reqErr.Code == maelstrom.KeyDoesNotExist {
			v = 0
		} else {
			return err
		}
	}

	total := v
	for _, id := range s.node.NodeIDs() {
		if id == s.node.ID() {
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		v, err := s.kv.ReadInt(ctx, id)
		if err != nil {
			reqErr := &maelstrom.RPCError{}
			if errors.As(err, &reqErr) && reqErr.Code == maelstrom.KeyDoesNotExist {
				v = 0
			} else {
				return err
			}
		}
		total += v
	}

	respBody := &ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Value:       total,
	}

	return s.node.Reply(msg, respBody)
}
