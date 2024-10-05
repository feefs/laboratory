package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type CommitOffsetsReqBody struct {
	maelstrom.MessageBody
	Offsets map[string]int `json:"offsets"`
}

func (s *server) CommitOffsetsHandler(msg maelstrom.Message) error {
	reqBody := &CommitOffsetsReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	s.committedOffsetsMu.Lock()
	for key, offset := range reqBody.Offsets {
		s.committedOffsets[key] = offset
	}
	s.committedOffsetsMu.Unlock()

	return s.node.Reply(msg, maelstrom.MessageBody{Type: "commit_offsets_ok"})
}
