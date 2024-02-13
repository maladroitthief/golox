package main

import (
	"fmt"
	"os"

	"github.com/maladroitthief/golox/main/pkg/lox"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	}
	l := lox.New()

	if len(args) == 1 {
		l.RunFile(args[0])
	} else {
		l.RunPrompt()
	}
}
