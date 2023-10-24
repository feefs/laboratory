package main

import (
	"broadcast/server"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	server := server.NewServer(node)

	node.Handle("broadcast", server.BroadcastHandler)
	node.Handle("read", server.ReadHandler)
	node.Handle("topology", server.TopologyHandler)

	if err := node.Run(); err != nil {
		panic(err)
	}
}
