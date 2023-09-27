package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Internal state
type state struct {
	Messages []int64  `json:"messages,omitempty"`
	Topology topology `json:"topology,omitempty"`
}
type clusterNode string
type topology map[string][]clusterNode

// Handler types
type broadcastReqBody struct {
	maelstrom.MessageBody
	Message int64 `json:"message,omitempty"`
}
type broadcastRespBody struct {
	maelstrom.MessageBody
}

type readReqBody struct {
	maelstrom.MessageBody
}
type readRespBody struct {
	maelstrom.MessageBody
	Messages []int64 `json:"messages,omitempty"`
}

type topologyReqBody struct {
	maelstrom.MessageBody
	Topology topology `json:"topology,omitempty"`
}
type topologyRespBody struct {
	maelstrom.MessageBody
}

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

func main() {
	node := maelstrom.NewNode()

	data := &state{}

	node.Handle("broadcast", broadcastHandler(node, data))
	node.Handle("read", readHandler(node, data))
	node.Handle("topology", topologyHandler(node, data))

	if err := node.Run(); err != nil {
		panic(err)
	}
}
