package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
}

func serve(_ context.Context, _ *app) error {
	fmt.Println("Hello 世界!")

	return nil
}
