package handlers

import (
	"broadcast/types"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Read(node *maelstrom.Node, nodeState *types.State) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		respBody := &types.ReadRespBody{
			MessageBody: maelstrom.MessageBody{Type: "read_ok"},
			Messages:    nodeState.Messages,
		}

		return node.Reply(msg, respBody)
	}
}
