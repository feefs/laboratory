package main

import (
	"broadcast/server"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	server := server.NewServer(node)

	node.Handle("init", server.InitHandler)
	node.Handle("topology", server.TopologyHandler)
	node.Handle("broadcast", server.BroadcastHandler)
	node.Handle("propagate", server.PropagateHandler)
	node.Handle("read", server.ReadHandler)

	if err := node.Run(); err != nil {
		panic(err)
	}
}
