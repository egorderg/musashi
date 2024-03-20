package internal

import (
	"strings"
	"testing"
)

func BenchmarkParser(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		parser := newParser(strings.NewReader("(print \"Hello World\")"))
		parser.parse()
	}
}

func TestParser(t *testing.T) {
	parser := newParser(strings.NewReader("(print \"Hello World\")"))
	program, _ := parser.parse()
	want := AstProgram{forms: []AstForm{
		{items: []AstItem{
			AstSymbol{name: "print"},
			AstDatum{id: datumString, value: "Hello World"},
		}},
	}}

	for e := range parser.errors {
		t.Fatalf("%v", e)
	}

	assertItems(t, program.forms[0].items, want.forms[0].items)
}

func assertItems(t *testing.T, l []AstItem, r []AstItem) {
	t.Helper()

	if len(l) != len(r) {
		t.Fatalf("length")
	}

	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			t.Fatalf("compare")
		}
	}
}
