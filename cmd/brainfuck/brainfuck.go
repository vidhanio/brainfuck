package main

import (
	"errors"
	"flag"
	"os"

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
		content, err := os.ReadFile(*file)
		if err != nil {
			panic(err)
		}

		code = string(content)
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
