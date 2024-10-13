package main

import (
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type GenerateRespBody struct {
	maelstrom.MessageBody
	Id int `json:"id"`
}

func main() {
	node := maelstrom.NewNode()

	ids := make(chan int) // unbuffered channel
	generating := false
	node.Handle("init", func(msg maelstrom.Message) error {
		if generating {
			return nil
		}
		intIdStr := node.ID()[1:]
		intId, err := strconv.Atoi(intIdStr)
		if err != nil {
			return err
		}
		go func() {
			for {
				ids <- intId
				intId += len(node.NodeIDs())
			}
		}()
		generating = true
		return node.Reply(msg, maelstrom.MessageBody{Type: "init_ok"})
	})

	node.Handle("generate", func(msg maelstrom.Message) error {
		return node.Reply(msg, &GenerateRespBody{maelstrom.MessageBody{Type: "generate_ok"}, <-ids})
	})

	if err := node.Run(); err != nil {
		panic(err)
	}
}
