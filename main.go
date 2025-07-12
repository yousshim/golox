package main

import (
	"bufio"
	"fmt"
	"os"

	"golox/scanner"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		println("Usage: golox [script]")
		os.Exit(42)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = run(string(bytes))
	if err != nil {
		os.Exit(65)
	}
	return nil
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		_ = run(string(line))
	}
}

func run(script string) error {
	tokens, err := scanner.Scan(script)
	if err != nil {
		return err
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}
