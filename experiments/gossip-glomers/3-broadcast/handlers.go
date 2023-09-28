package main

import (
	"encoding/json"

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

		respBody := &broadcastRespBody{}
		respBody.Type = "broadcast_ok"

		return node.Reply(msg, respBody)
	}
}

func readHandler(node *maelstrom.Node, nodeState *state) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &readReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		respBody := &readRespBody{}
		respBody.Messages = nodeState.Messages
		respBody.Type = "read_ok"

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

		respBody := &topologyRespBody{}
		respBody.Type = "topology_ok"

		return node.Reply(msg, respBody)
	}
}
