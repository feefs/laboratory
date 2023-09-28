package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

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
