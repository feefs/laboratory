package main

import (
	"kafka/server"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	server := server.NewServer(node)

	node.Handle("send", server.SendHandler)
	node.Handle("poll", server.PollHandler)
	node.Handle("commit_offsets", server.CommitOffsetsHandler)
	node.Handle("list_committed_offsets", server.ListCommittedOffsetsHandler)

	if err := node.Run(); err != nil {
		panic(err)
	}
}
