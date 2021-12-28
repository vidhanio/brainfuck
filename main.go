package brainfuck

type Brainfuck struct {
	Cells        []byte
	instructions []rune
	instruction  int
	current      int
	operators    map[rune]func(*Brainfuck)
}

func New() *Brainfuck {
	b := new(Brainfuck)

	b.Cells = []byte{0}

	b.operators = make(map[rune]func(*Brainfuck))
	b.AddOperator('+', increment)
	b.AddOperator('-', decrement)
	b.AddOperator('<', previous)
	b.AddOperator('>', next)
	b.AddOperator('.', print)
	b.AddOperator(',', read)
	b.AddOperator('[', loop)
	b.AddOperator(']', endLoop)

	return b
}

func (b *Brainfuck) AddOperator(operator rune, fn func(*Brainfuck)) *Brainfuck {
	b.operators[operator] = fn

	return b
}

func (b *Brainfuck) AddInstructions(instructions string) *Brainfuck {
	for _, instruction := range instructions {
		if _, ok := b.operators[instruction]; ok {
			b.instructions = append(b.instructions, instruction)
		}
	}

	return b
}

func (b *Brainfuck) SetInstructions(instructions string) *Brainfuck {
	b.instructions = []rune{}

	return b.AddInstructions(instructions)
}

func (b *Brainfuck) Run() {
	for b.instruction < len(b.instructions) {
		b.operators[b.instructions[b.instruction]](b)
		b.instruction++
	}
}

func (b *Brainfuck) Increment() *Brainfuck {
	b.instructions = append(b.instructions, '+')

	return b
}

func (b *Brainfuck) Decrement() *Brainfuck {
	b.instructions = append(b.instructions, '-')

	return b
}

func (b *Brainfuck) Next() *Brainfuck {
	b.instructions = append(b.instructions, '>')

	return b
}

func (b *Brainfuck) Previous() *Brainfuck {
	b.instructions = append(b.instructions, '<')

	return b
}

func (b *Brainfuck) Print() *Brainfuck {
	b.instructions = append(b.instructions, '.')

	return b
}

func (b *Brainfuck) Read() *Brainfuck {
	b.instructions = append(b.instructions, ',')

	return b
}

func (b *Brainfuck) Loop() *Brainfuck {
	b.instructions = append(b.instructions, '[')

	return b
}

func (b *Brainfuck) EndLoop() *Brainfuck {
	b.instructions = append(b.instructions, ']')

	return b
}
