package main

import (
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

	// for _, tok := range tokens {
	// 	tok.PrintDebugLn()
	// }

	parser := parser.NewParser(tokens)
	ast := parser.ParseProgram()

	fmt.Printf("%+v\n", ast)
}
