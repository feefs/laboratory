package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type CommitOffsetsReqBody struct {
	maelstrom.MessageBody
	Offsets offsets `json:"offsets"`
}

func (s *server) CommitOffsetsHandler(msg maelstrom.Message) error {
	return ErrNotImplemented
}
