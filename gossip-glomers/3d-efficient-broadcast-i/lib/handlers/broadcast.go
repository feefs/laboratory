package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"broadcast/lib/handlers/types"
	"broadcast/lib/state"
	"broadcast/lib/util"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func BroadcastHandler(node *maelstrom.Node, nodeState *state.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) (err error) {
		defer func() {
			if err != nil {
				return
			}
			respBody := &types.BroadcastRespBody{MessageBody: maelstrom.MessageBody{Type: "broadcast_ok"}}
			err = node.Reply(msg, respBody)
		}()

		reqBody := &types.BroadcastReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		if reqBody.PropagationID != "" && nodeState.ContainsPropagation(reqBody.PropagationID) {
			return nil
		}

		propagateID, err := util.GeneratePropagateID()
		if err != nil {
			return err
		}

		nodeState.AppendMessage(reqBody.Message)
		nodeState.AddPropagation(propagateID)

		// optimization: don't broadcast back to "n0" if "n0" broadcasted to this node
		if msg.Src == "n0" {
			return nil
		}

		broadcastReq := &types.BroadcastReqBody{
			MessageBody:   maelstrom.MessageBody{Type: "broadcast"},
			Message:       reqBody.Message,
			PropagationID: propagateID,
		}
		for _, nid := range nodeState.Topology[node.ID()] {
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
					_, err := node.SyncRPC(ctx, nb, broadcastReq)
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
}
