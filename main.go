package main

import (
	"delphifmt/formatter"
	"delphifmt/lexer"
	"delphifmt/parser"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("tests/hello_world/input.pas")
	if err != nil {
		log.Fatal(err)
	}

	lexer := lexer.NewLexer(string(data))
	tokens := lexer.LexSrc()

	parser := parser.NewParser(tokens)
	ast := parser.ParseProgram()

	formatter := formatter.NewFormatter()
	fmt.Print(formatter.Format(&ast))
}
