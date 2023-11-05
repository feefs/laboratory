package server

import (
	"sync"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const rpcTimeout = 100 * time.Millisecond

type server struct {
	node *maelstrom.Node
	kv   *maelstrom.KV
	kvmu sync.Mutex
}

func NewServer(node *maelstrom.Node, kv *maelstrom.KV) *server {
	return &server{node: node, kv: kv}
}
