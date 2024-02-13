package lox

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maladroitthief/golox/main/pkg/scanner"
)

type Lox struct {
	hadError bool
}

func New() *Lox {
	return &Lox{}
}

func (l *Lox) RunFile(filePath string) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Unable to read file %v.\nError: %+v\n", filePath, err)
		os.Exit(64)
	}

	l.run(string(data))

	if l.hadError {
		os.Exit(65)
	}
}

func (l *Lox) RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		l.run(line)
		l.hadError = false
	}

	err := scanner.Err()
	if err != nil {
		fmt.Printf("Scanner error: %+v\n", err)
	}
}

func (l *Lox) run(source string) {
	scanner := scanner.NewScanner(source)
	tokens, err := scanner.ScanTokens()

	if err != nil {
		l.report(err)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (l *Lox) report(err error) {
	fmt.Println(err)
	l.hadError = true
}
