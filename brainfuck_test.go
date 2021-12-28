package brainfuck_test

import (
	"fmt"
	"strings"

	"github.com/vidhanio/brainfuck"
)

func Example() {
	bf := brainfuck.New()

	bf.SetInstructions("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	bf.Run()

	// Output: Hello World!
}

func ExampleBrainfuck_AddOperator() {
	bf := brainfuck.New()

	bf.AddOperator('~', func(b *brainfuck.Brainfuck) {
		b.Pointer = 0
	})
	bf.SetInstructions("+>>>>>~")
	bf.Run()

	fmt.Println(bf.Cells[bf.Pointer])

	// Output: 1
}

func ExampleBrainfuck_Increment() {
	bf := brainfuck.New()

	bf.SetInstructions("+")
	bf.Run()

	fmt.Println(bf.Cells[bf.Pointer])

	// Output: 1
}

func ExampleBrainfuck_Decrement() {
	bf := brainfuck.New()

	bf.SetInstructions("++-")
	bf.Run()

	fmt.Println(bf.Cells[0])

	// Output: 1
}

func ExampleBrainfuck_Next() {
	bf := brainfuck.New()

	bf.SetInstructions(">")
	bf.Run()

	fmt.Println(bf.Pointer)

	// Output: 1
}

func ExampleBrainfuck_Previous() {
	bf := brainfuck.New()

	bf.SetInstructions(">><")
	bf.Run()

	fmt.Println(bf.Pointer)

	// Output: 1
}

func ExampleBrainfuck_Print() {
	bf := brainfuck.New()

	bf.SetInstructions("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.")
	bf.Run()

	// Output: A
}

func ExampleBrainfuck_Read() {
	bf := brainfuck.New()

	bf.Reader = strings.NewReader("A")

	bf.SetInstructions(",.")
	bf.Run()

	// Output: A
}

func ExampleBrainfuck_StartLoop() {
	bf := brainfuck.New()

	bf.SetInstructions("+++++[>+++++++++++++<-]>.")
	bf.Run()

	// Output: A
}

func ExampleBrainfuck_EndLoop() {
	bf := brainfuck.New()

	bf.SetInstructions("+++++[>+++++++++++++<-]>.")
	bf.Run()

	// Output: A
}
