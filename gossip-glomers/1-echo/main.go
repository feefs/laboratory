package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type EchoReqBody struct {
	maelstrom.MessageBody
	Echo string `json:"echo"`
}
type EchoRespBody struct {
	maelstrom.MessageBody
	Echo string `json:"echo"`
}

func main() {
	node := maelstrom.NewNode()

	node.Handle("echo", func(msg maelstrom.Message) error {
		reqBody := &EchoReqBody{}
		if err := json.Unmarshal(msg.Body, reqBody); err != nil {
			return err
		}
		return node.Reply(msg, &EchoRespBody{maelstrom.MessageBody{Type: "echo_ok"}, reqBody.Echo})
	})

	if err := node.Run(); err != nil {
		panic(err)
	}
}
