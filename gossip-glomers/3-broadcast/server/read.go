package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type ReadRespBody struct {
	maelstrom.MessageBody
	Messages []int `json:"messages"`
}

func (s *server) ReadHandler(msg maelstrom.Message) error {
	s.prepareReadMessagesChan <- struct{}{}

	respBody := &ReadRespBody{
		MessageBody: maelstrom.MessageBody{Type: "read_ok"},
		Messages:    <-s.readMessagesChan,
	}

	return s.node.Reply(msg, respBody)
}
