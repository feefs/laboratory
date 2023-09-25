package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ServiceWeaver/weaver"
)

type StrManipulator interface {
	Capitalize(ctx context.Context, input string) (string, error)
	Reverse(ctx context.Context, input string) (string, error)
}

type strManipulator struct {
	weaver.Implements[StrManipulator]
	weaver.WithConfig[strManipulatorConfig]
}

type strManipulatorConfig struct {
	InitMessage string `toml:"init_message"`
}

func (s *strManipulator) Init(_ context.Context) error {
	fmt.Printf("Initializing strManipulator with init_message=%s\n", s.Config().InitMessage)
	return nil
}

var _ weaver.NotRetriable = StrManipulator.Reverse

func (s *strManipulator) Capitalize(_ context.Context, input string) (string, error) {
	return strings.ToUpper(input), nil
}

func (s *strManipulator) Reverse(_ context.Context, input string) (string, error) {
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i += 1 {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
