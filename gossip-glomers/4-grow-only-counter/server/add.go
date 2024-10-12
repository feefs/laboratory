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

	errChan := make(chan error)
	go func() {
		var doneErr error
		for {
			ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
			defer cancel()
			currentValue, err := s.kv.ReadInt(ctx, s.node.ID())
			nextValue := reqBody.Delta
			if err == nil {
				nextValue += currentValue
			} else if maelstrom.ErrorCode(err) != maelstrom.KeyDoesNotExist {
				doneErr = err
				break
			}

			ctx, cancel = context.WithTimeout(context.Background(), rpcTimeout)
			defer cancel()
			err = s.kv.CompareAndSwap(ctx, s.node.ID(), currentValue, nextValue, true)
			if err == nil {
				break
			} else if maelstrom.ErrorCode(err) != maelstrom.PreconditionFailed {
				doneErr = err
				break
			}
		}
		errChan <- doneErr
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "add_ok"})
}
