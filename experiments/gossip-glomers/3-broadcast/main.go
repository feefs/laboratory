package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	nodeState := &state{
		Messages:   make([]int64, 0),
		Topology:   make(topology),
		Propagated: make(map[propagateID]struct{}),
	}

	node.Handle("broadcast", broadcastHandler(node, nodeState))
	node.Handle("propagate", propagateHandler(node, nodeState))
	node.Handle("read", readHandler(node, nodeState))
	node.Handle("topology", topologyHandler(node, nodeState))

	if err := node.Run(); err != nil {
		panic(err)
	}
}
