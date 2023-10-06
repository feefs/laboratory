package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Body struct {
	maelstrom.MessageBody
	Echo string `json:"echo,omitempty"`
}

func main() {
	node := maelstrom.NewNode()

	node.Handle("echo", func(msg maelstrom.Message) error {
		body := &Body{}
		if err := json.Unmarshal(msg.Body, body); err != nil {
			return err
		}

		body.Type = "echo_ok"

		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		panic(err)
	}
}
