package main

import (
	"counter/server"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	kv := maelstrom.NewSeqKV(node)
	server := server.NewServer(node, kv)

	node.Handle("add", server.AddHandler)
	node.Handle("read", server.ReadHandler)

	if err := node.Run(); err != nil {
		panic(err)
	}
}
