package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type PropagateReqBody struct {
	maelstrom.MessageBody
	Messages []int `json:"messages"`
}

func (s *server) PropagateHandler(msg maelstrom.Message) error {
	reqBody := &PropagateReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	for _, msg := range reqBody.Messages {
		s.mp.messagesChan <- msg
	}

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "propagate_ok"})
}
