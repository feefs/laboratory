package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReqBody struct {
	maelstrom.MessageBody
	Message int `json:"message"`
}

func (s *server) BroadcastHandler(msg maelstrom.Message) error {
	reqBody := &BroadcastReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	s.messagesChan <- reqBody.Message

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
}
