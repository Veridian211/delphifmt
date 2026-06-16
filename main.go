package main

import (
	"delphifmt/formatter"
	"delphifmt/lexer"
	"delphifmt/parser"
	"delphifmt/token"
	"fmt"
	"os"
)

func main() {
	filename := "./formatter/tests/comments/input.pas"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: File %s not found.\n", filename)
		os.Exit(1)
	}

	lexer := lexer.NewLexer(string(data))
	tokens := lexer.LexSrc()

	token.PrintDebugLn(tokens)

	parser := parser.NewParser(tokens)
	ast, ok := parser.ParseProgram()
	if !ok {
		for _, err := range parser.GetErrors() {
			fmt.Println(err.String())
		}
		os.Exit(1)
	}

	formatter := formatter.NewFormatter()
	fmt.Print(formatter.Format(&ast))
}
