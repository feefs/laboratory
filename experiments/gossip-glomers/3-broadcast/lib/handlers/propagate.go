package handlers

import (
	"context"
	"encoding/json"
	"errors"

	"broadcast/types"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func PropagateHandler(node *maelstrom.Node, nodeState *types.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &types.PropagateReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		respBody := &types.PropagateRespBody{MessageBody: maelstrom.MessageBody{Type: "propagate_ok"}}

		if _, ok := nodeState.Propagated[reqBody.PropagateID]; ok {
			return node.Reply(msg, respBody)
		}
		nodeState.Propagated[reqBody.PropagateID] = struct{}{}
		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		errs := []error{}
		propagateReq := &types.PropagateReqBody{
			MessageBody: maelstrom.MessageBody{Type: "propagate"},
			Message:     reqBody.Message,
			PropagateID: reqBody.PropagateID,
		}
		for _, neighbor := range nodeState.Topology[node.ID()] {
			if _, err := node.SyncRPC(context.Background(), neighbor, propagateReq); err != nil {
				errs = append(errs, err)
			}
		}
		err := errors.Join(errs...)

		if err != nil {
			respBody.Code = maelstrom.Crash
			respBody.Text = err.Error()
		}

		return node.Reply(msg, respBody)
	}
}
