package internal

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	parser := newParser(strings.NewReader("(print \"Hello World\")"))
	program, _ := parser.parse()

	t.Logf("%v", program)

	for e := range parser.errors {
		t.Fatalf("%v", e)
	}
}
