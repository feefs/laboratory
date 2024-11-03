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

	if msg.Src[0] == 'c' {
		for _, id := range s.node.NodeIDs() {
			if id == s.node.ID() {
				continue
			}
			go func() {
				for {
					ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
					defer cancel()
					_, err := s.node.SyncRPC(ctx, id, &BroadcastReqBody{
						MessageBody: maelstrom.MessageBody{Type: "broadcast"},
						Message:     reqBody.Message,
					})
					if err == nil {
						break
					}
				}
			}()
		}
	}

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
}
