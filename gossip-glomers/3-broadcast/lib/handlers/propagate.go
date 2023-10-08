package handlers

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"broadcast/lib/state"
	"broadcast/types"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func PropagateHandler(node *maelstrom.Node, nodeState *state.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &types.PropagateReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		respBody := &types.PropagateRespBody{MessageBody: maelstrom.MessageBody{Type: "propagate_ok"}}

		if nodeState.Propagation.SyncContains(reqBody.PropagateID) {
			return node.Reply(msg, respBody)
		}
		nodeState.Propagation.SyncAdd(reqBody.PropagateID)

		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		propagateReq := &types.PropagateReqBody{
			MessageBody: maelstrom.MessageBody{Type: "propagate"},
			Message:     reqBody.Message,
			PropagateID: reqBody.PropagateID,
		}
		wg := sync.WaitGroup{}
		for _, neighbor := range nodeState.Topology[node.ID()] {
			wg.Add(1)
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
				wg.Done()
			}()
		}

		wg.Wait()

		return node.Reply(msg, respBody)
	}
}
