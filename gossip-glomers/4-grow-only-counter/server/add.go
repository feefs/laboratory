package server

import (
	"context"
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type AddReqBody struct {
	maelstrom.MessageBody
	Delta int `json:"delta"`
}

func (s *server) AddHandler(msg maelstrom.Message) error {
	reqBody := &AddReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	// lock to prevent multiple add requests/goroutines in this node from overwriting a value
	s.kvmu.Lock()
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	value, err := s.kv.ReadInt(ctx, s.node.ID())
	if err != nil {
		return err
	}

	// uncomment the following line and remove the surrounding mutex calls to induce the race condition
	// time.Sleep(500 * time.Millisecond)

	ctx, cancel = context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	err = s.kv.Write(ctx, s.node.ID(), value+reqBody.Delta)
	if err != nil {
		return err
	}
	s.kvmu.Unlock()

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "add_ok"})
}
