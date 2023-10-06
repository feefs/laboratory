package handlers

import (
	"context"
	"encoding/json"
	"errors"

	"broadcast/lib/util"
	"broadcast/types"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func BroadcastHandler(node *maelstrom.Node, nodeState *types.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &types.BroadcastReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		respBody := &types.BroadcastRespBody{MessageBody: maelstrom.MessageBody{Type: "broadcast_ok"}}

		propagateID, err := util.GeneratePropagateID()
		if err != nil {
			respBody.Code = maelstrom.Crash
			respBody.Text = err.Error()
			return node.Reply(msg, respBody)
		}

		nodeState.Propagated[propagateID] = struct{}{}

		errs := []error{}
		propagateReq := &types.PropagateReqBody{
			MessageBody: maelstrom.MessageBody{Type: "propagate"},
			Message:     reqBody.Message,
			PropagateID: propagateID,
		}
		for _, neighbor := range nodeState.Topology[node.ID()] {
			if _, err := node.SyncRPC(context.Background(), neighbor, propagateReq); err != nil {
				errs = append(errs, err)
			}
		}

		if err := errors.Join(errs...); err != nil {
			respBody.Code = maelstrom.Crash
			respBody.Text = err.Error()
		}

		return node.Reply(msg, respBody)
	}
}
