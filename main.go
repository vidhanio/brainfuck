package brainfuck

import (
	"bufio"
	"fmt"
	"os"
)

type Brainfuck struct {
	Instructions []rune  // The instructions to execute
	Instruction  int     // The current instruction
	Cells        []uint8 // Cells
	Current      int     // The current cell
}

func New(instructions string) *Brainfuck {
	var instructs []rune

	for _, instruction := range instructions {
		if instruction == '+' || instruction == '-' || instruction == '<' || instruction == '>' || instruction == '.' || instruction == ',' || instruction == '[' || instruction == ']' {
			instructs = append(instructs, instruction)
		}
	}

	return &Brainfuck{
		Instructions: instructs,
		Instruction:  0,
		Cells:        []uint8{0},
		Current:      0,
	}
}

func (b *Brainfuck) Run() {
	for b.Instruction < len(b.Instructions) {
		switch b.Instructions[b.Instruction] {
		case '+':
			b.Increment()
		case '-':
			b.Decrement()
		case '<':
			b.Previous()
		case '>':
			b.Next()
		case '.':
			b.Print()
		case ',':
			b.Read()
		case '[':
			b.Loop()
		case ']':
			b.EndLoop()
		}

		b.Instruction++
	}
}

func (b *Brainfuck) Next() *Brainfuck {
	b.Current++
	if b.Current >= len(b.Cells) {
		b.Cells = append(b.Cells, 0)
	}

	return b
}

func (b *Brainfuck) Previous() *Brainfuck {
	if b.Current > 0 {
		b.Current--
	}

	return b
}

func (b *Brainfuck) Increment() *Brainfuck {
	b.Cells[b.Current]++

	return b
}

func (b *Brainfuck) Decrement() *Brainfuck {
	b.Cells[b.Current]--

	return b
}

func (b *Brainfuck) Print() *Brainfuck {
	fmt.Print(string(b.Cells[b.Current]))

	return b
}

func (b *Brainfuck) Read() *Brainfuck {
	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()
	b.Cells[b.Current] = uint8(char)

	return b
}

func (b *Brainfuck) Loop() *Brainfuck {
	if b.Cells[b.Current] == 0 {
		b.Instruction = findMatching(1, b.Instructions, b.Instruction)
		return b
	}

	return b
}

func (b *Brainfuck) EndLoop() *Brainfuck {
	if b.Cells[b.Current] != 0 {
		b.Instruction = findMatching(-1, b.Instructions, b.Instruction)
		return b
	}

	return b
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
