package main

import (
	"os"
	"testing"
)

// TestMainWithName runs the main function with the "--name Bill" argument
func TestMainWithName(t *testing.T) {
	// Simulate the command line input
	os.Args = []string{"helloapp", "--name", "Bill"}
	main()
}
