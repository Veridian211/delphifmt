package formatter_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"delphifmt/formatter"
	"delphifmt/lexer"
	"delphifmt/parser"
)

func TestGolden(t *testing.T) {
	cases, err := os.ReadDir("./tests")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		if !c.IsDir() {
			continue
		}
		t.Run(c.Name(), func(t *testing.T) {
			dir := filepath.Join("./tests", c.Name())
			input, err := os.ReadFile(filepath.Join(dir, "input.pas"))
			if err != nil {
				t.Fatal(err)
			}
			expected, err := os.ReadFile(filepath.Join(dir, "expected.pas"))
			if err != nil {
				t.Fatal(err)
			}

			output := formatSource(string(input))

			outputPath := filepath.Join(dir, "output.pas")
			if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
				t.Fatalf("failed to write output.pas: %v", err)
			}

			if output != string(expected) {
				diff, _ := exec.Command("git", "diff", "--no-index",
					filepath.Join(dir, "expected.pas"),
					outputPath,
				).Output()
				t.Errorf("output mismatch:\n%s", diff)
			}

			idempotentOutput := formatSource(string(expected))

			idempotentOutputPath := filepath.Join(dir, "output_idempotent.pas")
			if err := os.WriteFile(idempotentOutputPath, []byte(idempotentOutput), 0644); err != nil {
				t.Fatalf("failed to write output_idempotent.pas: %v", err)
			}

			if idempotentOutput != string(expected) {
				diff, _ := exec.Command("git", "diff", "--no-index",
					filepath.Join(dir, "expected.pas"),
					idempotentOutputPath,
				).Output()
				t.Errorf("formatter not idempotent:\n%s", diff)
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
