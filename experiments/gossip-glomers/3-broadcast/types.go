package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Internal state
type state struct {
	Messages []int64  `json:"messages,omitempty"`
	Topology topology `json:"topology,omitempty"`
}
type clusterNode string
type topology map[string][]clusterNode

// Handlers
type broadcastReqBody struct {
	maelstrom.MessageBody
	Message int64 `json:"message,omitempty"`
}
type broadcastRespBody struct {
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
