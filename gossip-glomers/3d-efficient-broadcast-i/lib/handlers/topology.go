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

		nodes := []string{}
		if node.ID() == "n0" {
			for _, nid := range node.NodeIDs() {
				if nid == "n0" {
					continue
				}
				nodes = append(nodes, nid)
			}
		} else {
			nodes = append(nodes, "n0")
		}
		nodeState.Topology = state.Topology{node.ID(): nodes}

		respBody := &types.TopologyRespBody{MessageBody: maelstrom.MessageBody{Type: "topology_ok"}}

		return node.Reply(msg, respBody)
	}
}
