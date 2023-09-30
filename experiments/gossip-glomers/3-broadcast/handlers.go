package main

import (
	"context"
	"encoding/json"
	"errors"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Handlers
func broadcastHandler(node *maelstrom.Node, nodeState *state) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &broadcastReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		respBody := &broadcastRespBody{maelstrom.MessageBody{Type: "broadcast_ok"}}

		propagateID, err := generatePropagateID()
		if err != nil {
			respBody.Code = maelstrom.Crash
			respBody.Text = err.Error()
			return node.Reply(msg, respBody)
		}

		nodeState.Propagated[propagateID] = struct{}{}

		errs := []error{}
		propagateReq := &propagateReqBody{
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

func propagateHandler(node *maelstrom.Node, nodeState *state) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &propagateReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		respBody := &propagateRespBody{maelstrom.MessageBody{Type: "propagate_ok"}}

		if _, ok := nodeState.Propagated[reqBody.PropagateID]; ok {
			return node.Reply(msg, respBody)
		}
		nodeState.Propagated[reqBody.PropagateID] = struct{}{}
		nodeState.Messages = append(nodeState.Messages, reqBody.Message)

		errs := []error{}
		propagateReq := &propagateReqBody{
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

func readHandler(node *maelstrom.Node, nodeState *state) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		respBody := &readRespBody{
			MessageBody: maelstrom.MessageBody{Type: "read_ok"},
			Messages:    nodeState.Messages,
		}

		return node.Reply(msg, respBody)
	}
}

func topologyHandler(node *maelstrom.Node, nodeState *state) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &topologyReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		nodeState.Topology = reqBody.Topology

		respBody := &topologyRespBody{maelstrom.MessageBody{Type: "topology_ok"}}

		return node.Reply(msg, respBody)
	}
}
