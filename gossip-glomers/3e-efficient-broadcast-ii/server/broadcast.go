package server

import (
	"broadcast/server/rpc"
	"broadcast/server/state"
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReqBody struct {
	maelstrom.MessageBody
	BroadcastID state.BroadcastID `json:"broadcast_id"` // Non-empty if coming from a child node
	Message     int64             `json:"message"`
}

func (s *Server) BroadcastHandler(msg maelstrom.Message) (err error) {
	defer func() {
		if err != nil {
			return
		}
		err = s.Node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
	}()

	reqBody := &BroadcastReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	if s.Node.ID() != "n0" {
		return s.handleBroadcastChild(reqBody)
	} else {
		return s.handleBroadcastParent(reqBody)
	}
}

func (s *Server) handleBroadcastChild(reqBody *BroadcastReqBody) error {
	id, err := state.GenerateBroadcastID()
	if err != nil {
		return err
	}

	broadcastReq := &BroadcastReqBody{
		MessageBody: maelstrom.MessageBody{Type: "broadcast"},
		BroadcastID: id,
		Message:     reqBody.Message,
	}
	go rpc.Retry(s.Node, "n0", broadcastReq)

	return nil
}

func (s *Server) handleBroadcastParent(reqBody *BroadcastReqBody) error {
	if reqBody.BroadcastID != "" {
		if s.State.HasBroadcastID(reqBody.BroadcastID) {
			return nil
		}
		s.State.AddBroadcastID(reqBody.BroadcastID)
	}

	s.State.AppendMessages(reqBody.Message)

	s.State.Batch.Input <- reqBody.Message

	return nil
}
