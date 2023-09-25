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
}

func (s *strManipulator) Init(_ context.Context) error {
	fmt.Println("Initialized strManipulator!")
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
