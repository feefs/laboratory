package main

import (
	"fmt"
	"np/cmd"
	"os"

	"github.com/fatih/color"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
