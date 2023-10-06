package types

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Internal state
type Topology map[string][]string
type PropagateID string

type State struct {
	Messages   []int64  `json:"messages,omitempty"`
	Topology   Topology `json:"topology,omitempty"`
	Propagated map[PropagateID]struct{}
}

// Handlers
type BroadcastReqBody struct {
	maelstrom.MessageBody
	Message int64 `json:"message,omitempty"`
}
type BroadcastRespBody struct {
	maelstrom.MessageBody
}

type PropagateReqBody struct {
	maelstrom.MessageBody
	PropagateID PropagateID `json:"propagate_id,omitempty"`
	Message     int64       `json:"message,omitempty"`
}
type PropagateRespBody struct {
	maelstrom.MessageBody
}

type ReadRespBody struct {
	maelstrom.MessageBody
	Messages []int64 `json:"messages"`
}

type TopologyReqBody struct {
	maelstrom.MessageBody
	Topology Topology `json:"topology,omitempty"`
}
type TopologyRespBody struct {
	maelstrom.MessageBody
}
