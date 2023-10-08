package types

import (
	"broadcast/lib/state"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

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
	PropagationID state.PropagationID `json:"propagation_id,omitempty"`
	Message       int64               `json:"message,omitempty"`
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
	Topology state.Topology `json:"topology,omitempty"`
}
type TopologyRespBody struct {
	maelstrom.MessageBody
}
