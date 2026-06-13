package formatter_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"delphifmt/formatter"
	"delphifmt/lexer"
	"delphifmt/parser"
)

func TestGolden(t *testing.T) {
	cases, err := os.ReadDir("../tests")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		if !c.IsDir() {
			continue
		}
		t.Run(c.Name(), func(t *testing.T) {
			dir := filepath.Join("../tests", c.Name())
			input, err := os.ReadFile(filepath.Join(dir, "input.pas"))
			if err != nil {
				t.Fatal(err)
			}
			expected, err := os.ReadFile(filepath.Join(dir, "expected.pas"))
			if err != nil {
				t.Fatal(err)
			}

			got := formatSource(string(input))
			if got != string(expected) {
				t.Errorf("output mismatch\ngot:\n%s\nwant:\n%s", got, string(expected))
			}

			got2 := formatSource(got)
			if got2 != got {
				t.Errorf("formatter not idempotent\nfirst pass:\n%s\nsecond pass:\n%s", got, got2)
			}
		})
	}
}

func formatSource(src string) string {
	tokens := lexer.NewLexer(src).LexSrc()

	parser := parser.NewParser(tokens)
	node, ok := parser.ParseProgram()
	if !ok {
		for _, err := range parser.GetErrors() {
			fmt.Println(err)
		}
	}
	return formatter.NewFormatter().Format(&node)
}
