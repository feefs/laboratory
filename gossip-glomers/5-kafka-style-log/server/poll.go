package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type PollReqBody struct {
	maelstrom.MessageBody
	Offsets map[string]int `json:"offsets"`
}
type PollRespBody struct {
	maelstrom.MessageBody
	Msgs map[string]([]([]int)) `json:"msgs"`
}

func (s *server) PollHandler(msg maelstrom.Message) error {
	reqBody := &PollReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	result := make(map[string]([]([]int)), len(s.offsets))

	s.offsetsmu.Lock()
	for key, offset := range reqBody.Offsets {
		if messages, ok := s.offsets[key]; ok {
			result[key] = []([]int){}
			for _, msg := range messages {
				if msg.offset < offset {
					continue
				}
				result[key] = append(result[key], []int{msg.offset, msg.value})
			}
		}
	}
	s.offsetsmu.Unlock()

	respBody := &PollRespBody{MessageBody: maelstrom.MessageBody{Type: "poll_ok"}, Msgs: result}

	return s.node.Reply(msg, respBody)
}
