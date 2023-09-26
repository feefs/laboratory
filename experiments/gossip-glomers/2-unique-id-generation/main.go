package main

import (
	"encoding/json"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Body struct {
	maelstrom.MessageBody
	Id int64 `json:"id"`
}

func main() {
	node := maelstrom.NewNode()

	ids := make(chan int64) // Unbuffered channel

	node.Handle("init", func(msg maelstrom.Message) error {
		intIdStr := node.ID()[1:]
		intId, err := strconv.ParseInt(intIdStr, 10, 64)
		if err != nil {
			return err
		}

		// send first int64 to the ids channel in a new goroutine so this function doesn't block
		go func() {
			ids <- intId
		}()

		return node.Reply(msg, maelstrom.MessageBody{Type: "init_ok"})
	})

	node.Handle("generate", func(msg maelstrom.Message) error {
		body := &Body{}
		if err := json.Unmarshal(msg.Body, body); err != nil {
			return err
		}

		id := <-ids
		// defer sending the next int64 to the ids channel so this function doesn't block
		defer func() {
			ids <- id + int64(len(node.NodeIDs()))
		}()

		body.Id = id
		body.Type = "generate_ok"

		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		panic(err)
	}
}
