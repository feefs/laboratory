package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Internal state
type topology map[string][]string
type propagateID string

type state struct {
	Messages   []int64  `json:"messages,omitempty"`
	Topology   topology `json:"topology,omitempty"`
	Propagated map[propagateID]struct{}
}

// Handlers
type broadcastReqBody struct {
	maelstrom.MessageBody
	Message int64 `json:"message,omitempty"`
}
type broadcastRespBody struct {
	maelstrom.MessageBody
}

type propagateReqBody struct {
	maelstrom.MessageBody
	PropagateID propagateID `json:"propagate_id,omitempty"`
	Message     int64       `json:"message,omitempty"`
}
type propagateRespBody struct {
	maelstrom.MessageBody
}

type readRespBody struct {
	maelstrom.MessageBody
	Messages []int64 `json:"messages"`
}

type topologyReqBody struct {
	maelstrom.MessageBody
	Topology topology `json:"topology,omitempty"`
}
type topologyRespBody struct {
	maelstrom.MessageBody
}
