package brainfuck

import (
	"bufio"
	"fmt"
)

func next(b *Brainfuck) {
	b.Pointer++
	if b.Pointer >= len(b.Cells) {
		b.Cells = append(b.Cells, 0)
	}
}

func previous(b *Brainfuck) {
	if b.Pointer > 0 {
		b.Pointer--
	}
}

func increment(b *Brainfuck) {
	b.Cells[b.Pointer]++
}

func decrement(b *Brainfuck) {
	b.Cells[b.Pointer]--
}

func print(b *Brainfuck) {
	fmt.Fprint(b.Writer, string(b.Cells[b.Pointer]))
}

func read(b *Brainfuck) {
	b.Cells[b.Pointer], _ = bufio.NewReader(b.Reader).ReadByte()
}

func startLoop(b *Brainfuck) {
	if b.Cells[b.Pointer] == 0 {
		b.Instruction = findMatching(1, b.Instructions, b.Instruction)
	}
}

func endLoop(b *Brainfuck) {
	if b.Cells[b.Pointer] != 0 {
		b.Instruction = findMatching(-1, b.Instructions, b.Instruction)
	}
}

func findMatching(direction int, instructions []rune, instruction int) int {
	counter := 0
	if direction == 1 {
		for i := instruction; i < len(instructions); i++ {
			if instructions[i] == ']' {
				counter++
			} else if instructions[i] == '[' {
				counter--
			}

			if counter == 0 {
				return i
			}
		}
	} else {
		for i := instruction; i >= 0; i-- {
			if instructions[i] == ']' {
				counter++
			} else if instructions[i] == '[' {
				counter--
			}

			if counter == 0 {
				return i
			}
		}
	}

	return -1
}
