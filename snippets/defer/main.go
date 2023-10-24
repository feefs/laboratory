package main

import (
	"errors"
	"fmt"
)

func foo(bar func() error) (ret int, err error) {
	defer func() {
		if err == nil {
			return
		}
		ret = -1
		fmt.Println("ret value modified")
	}()

	err = bar()

	return 100, err
}

func main() {
	bar1 := func() error {
		return nil
	}

	bar2 := func() error {
		return errors.New("error")
	}

	fmt.Println(foo(bar1))
	fmt.Println(foo(bar2))
}
