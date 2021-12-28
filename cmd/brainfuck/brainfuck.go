package main

import (
	"errors"
	"flag"

	"github.com/vidhanio/brainfuck"
)

func main() {
	// Intialize flags.
	file := flag.String("i", "", "brainfuck file to interpret")
	input := flag.String("s", "", "brainfuck string to interpret")
	flag.Parse()

	// Check if we have a file or string to interpret.
	var code string

	if *file != "" {
		code = *file
	} else if *input != "" {
		code = *input
	} else {
		panic(errors.New("no input provided"))
	}

	// Create a new brainfuck interpreter.
	bf := brainfuck.New()
	bf.SetInstructions(code)
	bf.Run()
}
