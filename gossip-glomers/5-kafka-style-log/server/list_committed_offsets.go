package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ListCommittedOffsetsReqBody struct {
	maelstrom.MessageBody
	Keys []string `json:"keys"`
}
type ListCommittedOffsetsRespBody struct {
	maelstrom.MessageBody
	Offsets map[string]int `json:"offsets"`
}

func (s *server) ListCommittedOffsetsHandler(msg maelstrom.Message) error {
	reqBody := &ListCommittedOffsetsReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	result := make(map[string]int)

	s.committedOffsetsMu.Lock()
	for _, key := range reqBody.Keys {
		if offset, ok := s.committedOffsets[key]; ok {
			result[key] = offset
		}
	}
	s.committedOffsetsMu.Unlock()

	respBody := &ListCommittedOffsetsRespBody{
		MessageBody: maelstrom.MessageBody{Type: "list_committed_offsets_ok"},
		Offsets:     result,
	}

	return s.node.Reply(msg, respBody)
}
