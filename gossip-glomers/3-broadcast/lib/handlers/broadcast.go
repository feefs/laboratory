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
	return func(msg maelstrom.Message) error {
		reqBody := &types.BroadcastReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		propagateID, err := util.GeneratePropagateID()
		if err != nil {
			return err
		}
		nodeState.Messages = append(nodeState.Messages, reqBody.Message)
		nodeState.AddPropagation(propagateID)

		// propagation goroutine
		go func() {
			propagateReq := &types.PropagateReqBody{
				MessageBody:   maelstrom.MessageBody{Type: "propagate"},
				Message:       reqBody.Message,
				PropagationID: propagateID,
			}
			for _, neighbor := range nodeState.Topology[node.ID()] {
				nb := neighbor
				go func() {
					wait := 100 * time.Millisecond
					attempts := 0
					attempt_limit := 10
					for attempts < attempt_limit {
						ctx, cancel := context.WithTimeout(context.Background(), wait)
						defer cancel()
						_, err := node.SyncRPC(ctx, nb, propagateReq)
						if err == nil {
							break
						}
						wait *= 2
						attempts += 1
					}
					if attempts == attempt_limit {
						log.Printf("Propagate timed out with %v attempts: propagateReq=%v neighbor=%v \n", attempt_limit, propagateReq, nb)
					}
				}()
			}
		}()

		respBody := &types.BroadcastRespBody{MessageBody: maelstrom.MessageBody{Type: "broadcast_ok"}}

		return node.Reply(msg, respBody)
	}
}
