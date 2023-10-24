package server

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ReadRespBody struct {
	maelstrom.MessageBody
	Messages []int64 `json:"messages"`
}

func (s *Server) ReadHandler(msg maelstrom.Message) (err error) {
	respBody := ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Messages:    s.state.ReadMessages(),
	}

	return s.node.Reply(msg, respBody)
}
