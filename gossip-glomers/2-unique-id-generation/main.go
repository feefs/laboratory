package main

import (
	"math/big"
	"math/rand/v2"
	"strconv"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type GenerateRespBody struct {
	maelstrom.MessageBody
	Id string `json:"id"`
}

func uuidV7(id int64) string {
	unixTsMs := big.NewInt(time.Now().UnixMilli())
	unixTsMs = unixTsMs.Lsh(unixTsMs, 80)

	ver := big.NewInt(7)
	ver = ver.Lsh(ver, 76)

	randA := big.NewInt(id)
	randA = randA.Lsh(randA, 64)

	variant := big.NewInt(2)
	variant = variant.Lsh(variant, 62)

	randB := big.NewInt(rand.Int64())
	randB = randB.And(randB, big.NewInt(0x3FFFFFFFFFFFFFFF))

	result := new(big.Int)
	result.Or(result, unixTsMs)
	result.Or(result, ver)
	result.Or(result, randA)
	result.Or(result, variant)
	result.Or(result, randB)

	return result.Text(16)
}

func main() {
	node := maelstrom.NewNode()

	var id int64
	node.Handle("init", func(msg maelstrom.Message) error {
		intIdStr := node.ID()[1:]
		parsedId, err := strconv.ParseInt(intIdStr, 10, 64)
		if err != nil {
			return err
		}
		id = parsedId
		return node.Reply(msg, maelstrom.MessageBody{Type: "init_ok"})
	})

	node.Handle("generate", func(msg maelstrom.Message) error {
		return node.Reply(msg, &GenerateRespBody{maelstrom.MessageBody{Type: "generate_ok"}, uuidV7(id)})
	})

	if err := node.Run(); err != nil {
		panic(err)
	}
}
