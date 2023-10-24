package main

import (
	"errors"
	"testing"
)

func TestFoo(t *testing.T) {
	bar1 := func() error {
		return nil
	}
	bar2 := func() error {
		return errors.New("error")
	}

	type test struct {
		bar    func() error
		ret    int
		nilerr bool
	}

	tests := []test{
		{
			bar:    bar1,
			ret:    100,
			nilerr: true,
		},
		{
			bar:    bar2,
			ret:    -1,
			nilerr: false,
		},
	}

	for _, test := range tests {
		ret, err := foo(test.bar)
		if ret != test.ret {
			t.Errorf("expected: %v, got: %v", test.ret, ret)
		}
		nilerr := err == nil
		if nilerr != test.nilerr {
			t.Errorf("expected: %v, got: %v", test.nilerr, nilerr)
		}
	}
}
