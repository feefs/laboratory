package server

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReqBody struct {
	maelstrom.MessageBody
	Message int `json:"message"`
}

func (s *server) BroadcastHandler(msg maelstrom.Message) error {
	if len(msg.Src) == 0 {
		return errors.New("empty caller type")
	}

	reqBody := &BroadcastReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	s.messagesChan <- reqBody.Message

	if s.node.ID() == "n0" {
		for _, id := range s.node.NodeIDs() {
			if id == "n0" || (msg.Src[0] == 'n' && id == msg.Src) {
				continue
			}
			go s.retryBroadcast(id, reqBody.Message)
		}
	} else {
		if msg.Src[0] == 'c' {
			go s.retryBroadcast("n0", reqBody.Message)
		}
	}

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
}

func (s *server) retryBroadcast(id string, message int) {
	timeout := 1 * time.Second
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_, err := s.node.SyncRPC(ctx, id, &BroadcastReqBody{
			MessageBody: maelstrom.MessageBody{Type: "broadcast"},
			Message:     message,
		})
		if err == nil {
			break
		}
		timeout += (250 * time.Millisecond)
	}
}
