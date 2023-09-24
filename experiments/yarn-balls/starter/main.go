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
	strManipulator weaver.Ref[StrManipulator]
}

func serve(ctx context.Context, app *app) error {
	input := "!界世 olleH"

	manipulator := app.strManipulator.Get()
	reversed, err := manipulator.Reverse(ctx, input)
	if err != nil {
		return err
	}

	capitalized, err := manipulator.Capitalize(ctx, reversed)
	if err != nil {
		return err
	}

	fmt.Println(capitalized)

	return nil
}
