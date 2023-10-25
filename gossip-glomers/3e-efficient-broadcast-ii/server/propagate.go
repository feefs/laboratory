package server

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type PropagateReqBody struct {
	maelstrom.MessageBody
	PropagationID PropagationID `json:"propagation_id"`
	Messages      []int64       `json:"messages"`
}

func (s *Server) PropagateHandler(msg maelstrom.Message) (err error) {
	defer func() {
		if err != nil {
			return
		}
		err = s.node.Reply(msg, &maelstrom.MessageBody{Type: "propagate_ok"})
	}()

	reqBody := &PropagateReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	if s.state.HasPropagationID(reqBody.PropagationID) {
		return nil
	}
	s.state.AddPropagationID(reqBody.PropagationID)

	s.state.AppendMessages(reqBody.Messages...)

	return nil
}
