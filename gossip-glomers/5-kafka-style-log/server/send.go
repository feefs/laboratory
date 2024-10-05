package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type SendReqBody struct {
	maelstrom.MessageBody
	Key string `json:"key"`
	Msg int    `json:"msg"`
}
type SendRespBody struct {
	maelstrom.MessageBody
	Offset int `json:"offset"`
}

func (s *server) SendHandler(msg maelstrom.Message) error {
	reqBody := &SendReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	storedOffset := -1
	s.offsetsmu.Lock()
	if messages, ok := s.offsets[reqBody.Key]; ok {
		lastMessage := messages[len(messages)-1]
		s.offsets[reqBody.Key] = append(messages, message{offset: lastMessage.offset + 1, value: reqBody.Msg})
		storedOffset = lastMessage.offset + 1
	} else {
		s.offsets[reqBody.Key] = []message{{offset: 1000, value: reqBody.Msg}}
		storedOffset = 1000
	}
	s.offsetsmu.Unlock()

	respBody := &SendRespBody{MessageBody: maelstrom.MessageBody{Type: "send_ok"}, Offset: storedOffset}

	return s.node.Reply(msg, respBody)
}
