package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReqBody struct {
	maelstrom.MessageBody
	PropagationID PropagationID `json:"propagation_id"`
	Message       int64         `json:"message"`
}
type BroadcastRespBody struct {
	maelstrom.MessageBody
}

func (s *Server) BroadcastHandler(msg maelstrom.Message) (err error) {
	defer func() {
		if err != nil {
			return
		}
		respBody := &BroadcastRespBody{MessageBody: maelstrom.MessageBody{Type: "broadcast_ok"}}
		err = s.node.Reply(msg, respBody)
	}()

	reqBody := &BroadcastReqBody{}
	if err := json.Unmarshal(msg.Body, reqBody); err != nil {
		return err
	}

	if reqBody.PropagationID != "" && s.state.ContainsPropagation(reqBody.PropagationID) {
		return nil
	}

	propagateID, err := GeneratePropagateID()
	if err != nil {
		return err
	}

	s.state.AppendMessage(reqBody.Message)
	s.state.AddPropagation(propagateID)

	// optimization: don't broadcast back to "n0" if "n0" broadcasted to this node
	if msg.Src == "n0" {
		return nil
	}

	broadcastReq := &BroadcastReqBody{
		MessageBody:   maelstrom.MessageBody{Type: "broadcast"},
		Message:       reqBody.Message,
		PropagationID: propagateID,
	}
	for _, nid := range s.state.Topology[s.node.ID()] {
		// optimization: don't broadcast to nid if nid broadcasted to this node
		if msg.Src == nid {
			continue
		}
		nb := nid
		go func() {
			timeout := 500 * time.Millisecond
			attempts := 0
			attempt_limit := 100
			for attempts < attempt_limit {
				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()
				_, err := s.node.SyncRPC(ctx, nb, broadcastReq)
				if err == nil {
					break
				}
				attempts += 1
				timeout += 500
			}
			if attempts == attempt_limit {
				log.Printf("Broadcast timed out with %v attempts: broadcastReq=%v neighbor=%v \n", attempt_limit, broadcastReq, nb)
			}
		}()
	}

	return nil
}
