package handlers

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"broadcast/lib/handlers/types"
	"broadcast/lib/state"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func PropagateHandler(node *maelstrom.Node, nodeState *state.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &types.PropagateReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		respBody := &types.PropagateRespBody{MessageBody: maelstrom.MessageBody{Type: "propagate_ok"}}

		if nodeState.ContainsPropagation(reqBody.PropagationID) {
			return node.Reply(msg, respBody)
		}
		nodeState.AddPropagation(reqBody.PropagationID)

		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		propagateReq := &types.PropagateReqBody{
			MessageBody:   maelstrom.MessageBody{Type: "propagate"},
			Message:       reqBody.Message,
			PropagationID: reqBody.PropagationID,
		}
		wg := sync.WaitGroup{}
		for _, neighbor := range nodeState.Topology[node.ID()] {
			wg.Add(1)
			nb := neighbor

			go func() {
				wait := 250 * time.Millisecond
				attempts := 0
				attempt_limit := 100
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
				wg.Done()
			}()
		}

		wg.Wait()

		return node.Reply(msg, respBody)
	}
}
