package rpc

import (
	"context"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Retry(node *maelstrom.Node, dest string, body any) {
	timeout := 500 * time.Millisecond
	attempts := 0
	attempt_limit := 100
	for attempts < attempt_limit {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_, err := node.SyncRPC(ctx, dest, body)
		if err == nil {
			break
		}
		attempts += 1
		timeout += 500
	}
	if attempts == attempt_limit {
		log.Printf("RPC timed out with %v attempts: dest=%v body=%v\n", attempts, dest, body)
	}
}
