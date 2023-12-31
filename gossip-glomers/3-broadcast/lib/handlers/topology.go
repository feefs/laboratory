package handlers

import (
	"encoding/json"

	"broadcast/lib/handlers/types"
	"broadcast/lib/state"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func TopologyHandler(node *maelstrom.Node, nodeState *state.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		reqBody := &types.TopologyReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}

		nodeState.Topology = reqBody.Topology

		respBody := &types.TopologyRespBody{MessageBody: maelstrom.MessageBody{Type: "topology_ok"}}

		return node.Reply(msg, respBody)
	}
}
