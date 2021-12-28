package main

import "github.com/vidhanio/brainfuck"

func main() {
	brainfuck.New().AddInstructions(`++++++++[>+>++++<<-]>++>>+<[-[>>+<<-]+>>]>+[
		-<<<[
			->[+[-]+>++>>>-<<]<[<]>>++++++[<<+++++>>-]+<<++.[-]<<
		]>.>+[>>]>+
	]`).Run()
}
