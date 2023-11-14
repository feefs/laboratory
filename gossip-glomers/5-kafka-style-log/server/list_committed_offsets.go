package server

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

type ListCommittedOffsetsReqBody struct {
	maelstrom.MessageBody
	Keys []string `json:"keys"`
}
type ListCommittedOffsetsRespBody struct {
	maelstrom.MessageBody
	Offsets offsets `json:"offsets"`
}

func (s *server) ListCommittedOffsetsHandler(msg maelstrom.Message) error {
	return ErrNotImplemented
}
