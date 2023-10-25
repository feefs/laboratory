package server

import (
	"broadcast/server/state"
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type PropagateReqBody struct {
	maelstrom.MessageBody
	PropagationID state.PropagationID `json:"propagation_id"`
	Messages      []int64             `json:"messages"`
}

func (s *Server) PropagateHandler(msg maelstrom.Message) (err error) {
	defer func() {
		if err != nil {
			return
		}
		err = s.Node.Reply(msg, &maelstrom.MessageBody{Type: "propagate_ok"})
	}()

	reqBody := &PropagateReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	if s.State.HasPropagationID(reqBody.PropagationID) {
		return nil
	}
	s.State.AddPropagationID(reqBody.PropagationID)

	s.State.AppendMessages(reqBody.Messages...)

	return nil
}
