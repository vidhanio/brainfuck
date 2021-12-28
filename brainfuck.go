package brainfuck

import (
	"io"
	"os"
)

/*
A Brainfuck interpreter.
*/
type Brainfuck struct {
	Cells        []byte                    // Slice of cells.
	Instructions []rune                    // Slice of instructions which will be run.
	Instruction  int                       // Index of the current instruction. (Not recommended to be changed manually.)
	Pointer      int                       // Index of the current cell. (Not recommended to be changed manually.)
	Operators    map[rune]func(*Brainfuck) // Map of operators.
	Reader       io.Reader                 // Reader to read bytes from.
	Writer       io.Writer                 // Writer to write bytes to.
}

/*
Create a new Brainfuck interpreter.

Automatically adds the following operators: + - > < . , [ ]
*/
func New() *Brainfuck {
	b := new(Brainfuck)

	b.Cells = []byte{0}
	b.Reader = os.Stdin
	b.Writer = os.Stdout

	b.Operators = make(map[rune]func(*Brainfuck))
	b.AddOperator('+', increment)
	b.AddOperator('-', decrement)
	b.AddOperator('<', previous)
	b.AddOperator('>', next)
	b.AddOperator('.', print)
	b.AddOperator(',', read)
	b.AddOperator('[', startLoop)
	b.AddOperator(']', endLoop)

	return b
}

/*
Add a new custom operator to the interpreter.
*/
func (b *Brainfuck) AddOperator(operator rune, fn func(*Brainfuck)) *Brainfuck {
	b.Operators[operator] = fn

	return b
}

/*
Set the instructions to run.
*/
func (b *Brainfuck) SetInstructions(instructions string) *Brainfuck {
	b.Instructions = []rune{}
	for _, instruction := range instructions {
		if _, ok := b.Operators[instruction]; ok {
			b.Instructions = append(b.Instructions, instruction)
		}
	}

	return b
}

/*
Run the instructions.
*/
func (b *Brainfuck) Run() {
	for b.Instruction < len(b.Instructions) {
		b.Operators[b.Instructions[b.Instruction]](b)
		b.Instruction++
	}
}

/*
Increment the value at the pointer.
*/
func (b *Brainfuck) Increment() *Brainfuck {
	b.Instructions = append(b.Instructions, '+')

	return b
}

/*
Decrement the value at the pointer.
*/
func (b *Brainfuck) Decrement() *Brainfuck {
	b.Instructions = append(b.Instructions, '-')

	return b
}

/*
Move the pointer to the next cell.
*/
func (b *Brainfuck) Next() *Brainfuck {
	b.Instructions = append(b.Instructions, '>')

	return b
}

/*
Move the pointer to the previous cell.
*/
func (b *Brainfuck) Previous() *Brainfuck {
	b.Instructions = append(b.Instructions, '<')

	return b
}

/*
Print the value at the pointer.
*/
func (b *Brainfuck) Print() *Brainfuck {
	b.Instructions = append(b.Instructions, '.')

	return b
}

/*
Read a character from stdin and store it at the pointer.
*/
func (b *Brainfuck) Read() *Brainfuck {
	b.Instructions = append(b.Instructions, ',')

	return b
}

/*
Start a loop.
*/
func (b *Brainfuck) StartLoop() *Brainfuck {
	b.Instructions = append(b.Instructions, '[')

	return b
}

/*
End a loop.
*/
func (b *Brainfuck) EndLoop() *Brainfuck {
	b.Instructions = append(b.Instructions, ']')

	return b
}
