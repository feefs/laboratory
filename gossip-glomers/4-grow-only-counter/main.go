package main

import maelstrom "github.com/jepsen-io/maelstrom/demo/go"

func main() {
	node := maelstrom.NewNode()

	if err := node.Run(); err != nil {
		panic(err)
	}
}
