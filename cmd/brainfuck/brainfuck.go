package main

import "github.com/vidhanio/brainfuck-go"

func main() {
	bf := brainfuck.New(`++++++++[>+>++++<<-]>++>>+<[-[>>+<<-]+>>]>+[
		-<<<[
			->[+[-]+>++>>>-<<]<[<]>>++++++[<<+++++>>-]+<<++.[-]<<
		]>.>+[>>]>+
	]`)
	bf.Run()
}
