package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

const BINARY = "np"

// Build the program
func Build() error {
	// Avoid fmt.Printf since this is a CLI application
	return sh.RunV("go", "build", "-o", BINARY)
}

// Clean artifacts
func Clean() error {
	fmt.Printf("rm %s\n", BINARY)
	return sh.Rm(BINARY)
}
