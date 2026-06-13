package parser

import (
	"fmt"
)

type ParseError struct {
	line    int
	col     int
	message string
}

func (err *ParseError) String() string {
	return fmt.Sprintf(
		"Line %d, Col %d: %s",
		err.line,
		err.col,
		err.message,
	)
}
