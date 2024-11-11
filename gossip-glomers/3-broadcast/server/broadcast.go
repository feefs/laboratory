package server

import (
	"encoding/json"
	"errors"

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

	s.mp.messagesChan <- reqBody.Message

	if s.node.ID() == "n0" {
		s.pp.propagationChan <- propagation{msg.Src, reqBody.Message}
	} else {
		if msg.Src[0] == 'c' {
			go s.resilientRpc("n0", &BroadcastReqBody{maelstrom.MessageBody{Type: "broadcast"}, reqBody.Message})
		}
	}

	return s.node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
}
