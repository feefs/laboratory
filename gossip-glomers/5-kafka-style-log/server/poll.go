package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type PollReqBody struct {
	maelstrom.MessageBody
	Offsets offsets `json:"offsets"`
}
type PollRespBody struct {
	maelstrom.MessageBody
	Msgs map[string]([]([]int))
}

func (s *server) PollHandler(msg maelstrom.Message) error {
	return ErrNotImplemented
}
