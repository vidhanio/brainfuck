package brainfuck

import (
	"bufio"
	"fmt"
	"os"
)

func next(b *Brainfuck) {
	b.current++
	if b.current >= len(b.Cells) {
		b.Cells = append(b.Cells, 0)
	}
}

func previous(b *Brainfuck) {
	if b.current > 0 {
		b.current--
	}
}

func increment(b *Brainfuck) {
	b.Cells[b.current]++
}

func decrement(b *Brainfuck) {
	b.Cells[b.current]--
}

func print(b *Brainfuck) {
	fmt.Print(string(b.Cells[b.current]))

}

func read(b *Brainfuck) {
	reader := bufio.NewReader(os.Stdin)
	char, _ := reader.ReadByte()
	b.Cells[b.current] = char
}

func loop(b *Brainfuck) {
	if b.Cells[b.current] == 0 {
		b.instruction = findMatching(1, b.instructions, b.instruction)
	}
}

func endLoop(b *Brainfuck) {
	if b.Cells[b.current] != 0 {
		b.instruction = findMatching(-1, b.instructions, b.instruction)
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
