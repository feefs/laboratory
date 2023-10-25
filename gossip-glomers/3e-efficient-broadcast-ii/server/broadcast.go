package server

import (
	"broadcast/server/rpc"
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReqBody struct {
	maelstrom.MessageBody
	BroadcastID BroadcastID `json:"broadcast_id"` // Non-empty if coming from a child node
	Message     int64       `json:"message"`
}

func (s *Server) BroadcastHandler(msg maelstrom.Message) (err error) {
	defer func() {
		if err != nil {
			return
		}
		err = s.node.Reply(msg, &maelstrom.MessageBody{Type: "broadcast_ok"})
	}()

	reqBody := &BroadcastReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	if s.node.ID() != "n0" {
		return s.handleBroadcastChild(reqBody)
	} else {
		return s.handleBroadcastParent(reqBody)
	}
}

func (s *Server) handleBroadcastChild(reqBody *BroadcastReqBody) error {
	id, err := GenerateBroadcastID()
	if err != nil {
		return err
	}

	broadcastReq := &BroadcastReqBody{
		MessageBody: maelstrom.MessageBody{Type: "broadcast"},
		BroadcastID: id,
		Message:     reqBody.Message,
	}
	go rpc.Retry(s.node, "n0", broadcastReq)

	return nil
}

func (s *Server) handleBroadcastParent(reqBody *BroadcastReqBody) error {
	if reqBody.BroadcastID != "" {
		if s.state.HasBroadcastID(reqBody.BroadcastID) {
			return nil
		}
		s.state.AddBroadcastID(reqBody.BroadcastID)
	}

	s.state.AppendMessages(reqBody.Message)

	s.state.batch.input <- reqBody.Message

	return nil
}
