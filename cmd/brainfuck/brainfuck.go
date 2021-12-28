package main

import (
	"github.com/vidhanio/brainfuck"
)

func main() {
	brainfuck.New().Increment().Increment().Increment().Increment().Increment().Increment().Increment().Increment().Loop().Next().Increment().Next().Increment().Increment().Increment().Increment().Previous().Previous().Decrement().EndLoop().Next().Increment().Increment().Next().Next().Increment().Previous().Loop().Decrement().Loop().Next().Next().Increment().Previous().Previous().Decrement().EndLoop().Increment().Next().Next().EndLoop().Next().Increment().Loop().Decrement().Previous().Previous().Previous().Loop().Decrement().Next().Loop().Increment().Loop().Decrement().EndLoop().Increment().Next().Increment().Increment().Next().Next().Next().Decrement().Previous().Previous().EndLoop().Previous().Loop().Previous().EndLoop().Next().Next().Increment().Increment().Increment().Increment().Increment().Increment().Loop().Previous().Previous().Increment().Increment().Increment().Increment().Increment().Next().Next().Decrement().EndLoop().Increment().Previous().Previous().Increment().Increment().Print().Loop().Decrement().EndLoop().Previous().Previous().EndLoop().Next().Print().Next().Increment().Loop().Next().Next().EndLoop().Next().Increment().EndLoop().Run()
}
