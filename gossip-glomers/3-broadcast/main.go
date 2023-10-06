package main

import (
	"broadcast/lib/handlers"
	"broadcast/types"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	nodeState := &types.State{
		Messages:   make([]int64, 0),
		Topology:   make(types.Topology),
		Propagated: make(map[types.PropagateID]struct{}),
	}

	node.Handle("broadcast", handlers.BroadcastHandler(node, nodeState))
	node.Handle("propagate", handlers.PropagateHandler(node, nodeState))
	node.Handle("read", handlers.Read(node, nodeState))
	node.Handle("topology", handlers.TopologyHandler(node, nodeState))

	if err := node.Run(); err != nil {
		panic(err)
	}
}
