package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type SendReqBody struct {
	maelstrom.MessageBody
	Key string `json:"key"`
	Msg string `json:"msg"`
}
type SendRespBody struct {
	maelstrom.MessageBody
	Offset int `json:"offset"`
}

func (s *server) SendHandler(msg maelstrom.Message) error {
	return ErrNotImplemented
}
