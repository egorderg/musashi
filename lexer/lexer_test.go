package lexer

import (
	"bufio"
	"strings"
	"testing"
)

type Dataset struct {
	name   string
	input  string
	tokens []Token
}

func TestLexer(t *testing.T) {
	datasets := []Dataset{
		commentData(),
		keywordData(),
		symbolData(),
		numberData(),
		stringData(),
		complexData(),
		concatData(),
		illegalData(),
	}

	for _, ds := range datasets {
		ds := ds

		t.Run(ds.name, func(t *testing.T) {
			t.Parallel()

			l := NewLexer(bufio.NewReader(strings.NewReader(ds.input)))
			tokens := make([]Token, 0, 10)

			for {
				tok := l.Next()
				if tok.Type == TokenEOF {
					break
				}

				tokens = append(tokens, tok)
			}

			if len(tokens) != len(ds.tokens) {
				for _, tok := range tokens {
					t.Logf("Lexer(%q) - %v", ds.input, tok)
				}

				t.Fatalf("Lexer(%q) - want %d tokens, got %d tokens", ds.input, len(ds.tokens), len(tokens))
			}

			for i := 0; i < len(tokens); i++ {
				if tokens[i] != ds.tokens[i] {
					t.Fatalf("tokenize(%q) - want %v, got %v", ds.input, ds.tokens[i], tokens[i])
				}
			}
		})
	}
}

func commentData() Dataset {
	return Dataset{name: "comments", input: "true#k平仮 2.234 dk jfk \n1.23#test", tokens: []Token{
		{TokenBool, "true", 0, 0},
		{TokenNumber, "1.23", 1, 0},
	}}
}

func keywordData() Dataset {
	return Dataset{name: "keywords", input: "true false nil", tokens: []Token{
		{TokenBool, "true", 0, 0},
		{TokenBool, "false", 0, 5},
		{TokenNil, nil, 0, 11},
	}}
}

func symbolData() Dataset {
	return Dataset{name: "symbols", input: "++\"str\"平032^t#test\n-a.bc :po\r\nset _ + * & | < > : % / -\n %-平仮-*k/", tokens: []Token{
		{TokenSymbol, "++", 0, 0},
		{TokenString, "str", 0, 2},
		{TokenSymbol, "平", 0, 7},
		{TokenNumber, "032", 0, 10},
		{TokenSymbol, "^t", 0, 13},
		{TokenSymbol, "-a.bc", 1, 0},
		{TokenSymbol, ":po", 1, 6},
		{TokenSymbol, "set", 2, 0},
		{TokenSymbol, "_", 2, 4},
		{TokenSymbol, "+", 2, 6},
		{TokenSymbol, "*", 2, 8},
		{TokenSymbol, "&", 2, 10},
		{TokenSymbol, "|", 2, 12},
		{TokenSymbol, "<", 2, 14},
		{TokenSymbol, ">", 2, 16},
		{TokenSymbol, ":", 2, 18},
		{TokenSymbol, "%", 2, 20},
		{TokenSymbol, "/", 2, 22},
		{TokenSymbol, "-", 2, 24},
		{TokenSymbol, "%-平仮-*k/", 3, 1},
	}}
}

func numberData() Dataset {
	return Dataset{name: "numbers", input: "05 7 -13\n47 -1.34\r\n27.234-b", tokens: []Token{
		{TokenNumber, "05", 0, 0},
		{TokenNumber, "7", 0, 3},
		{TokenSymbol, "-", 0, 5},
		{TokenNumber, "13", 0, 6},
		{TokenNumber, "47", 1, 0},
		{TokenSymbol, "-", 1, 3},
		{TokenNumber, "1.34", 1, 4},
		{TokenNumber, "27.234", 2, 0},
		{TokenSymbol, "-b", 2, 6},
	}}
}

func stringData() Dataset {
	return Dataset{name: "strings", input: "\"\" \"平仮名\"\"value\"\n \"long string\"", tokens: []Token{
		{TokenString, "", 0, 0},
		{TokenString, "平仮名", 0, 3},
		{TokenString, "value", 0, 14},
		{TokenString, "long string", 1, 1},
	}}
}

func complexData() Dataset {
	return Dataset{name: "complex", input: "#fn\r\n(+ 2 \n(!pos :pos-平 -2.7 \"hello 平\" 5\n))", tokens: []Token{
		{TokenLeftParen, nil, 1, 0},
		{TokenSymbol, "+", 1, 1},
		{TokenNumber, "2", 1, 3},
		{TokenLeftParen, nil, 2, 0},
		{TokenSymbol, "!pos", 2, 1},
		{TokenSymbol, ":pos-平", 2, 6},
		{TokenSymbol, "-", 2, 15},
		{TokenNumber, "2.7", 2, 16},
		{TokenString, "hello 平", 2, 20},
		{TokenNumber, "5", 2, 32},
		{TokenRightParen, nil, 3, 0},
		{TokenRightParen, nil, 3, 1},
	}}
}

func concatData() Dataset {
	return Dataset{name: "concat", input: "\"test 555\".34.23.\"\"平35-3;4\n4/!pos7true9\"test平\"\n#test\n\"test2\"", tokens: []Token{
		{TokenString, "test 555", 0, 0},
		{TokenSymbol, ".", 0, 10},
		{TokenNumber, "34.23", 0, 11},
		{TokenSymbol, ".", 0, 16},
		{TokenString, "", 0, 17},
		{TokenSymbol, "平", 0, 19},
		{TokenNumber, "35", 0, 22},
		{TokenSymbol, "-", 0, 24},
		{TokenNumber, "3", 0, 25},
		{TokenSymbol, ";", 0, 26},
		{TokenNumber, "4", 0, 27},
		{TokenNumber, "4", 1, 0},
		{TokenSymbol, "/!pos", 1, 1},
		{TokenNumber, "7", 1, 6},
		{TokenBool, "true", 1, 7},
		{TokenNumber, "9", 1, 11},
		{TokenString, "test平", 1, 12},
		{TokenString, "test2", 3, 0},
	}}
}

func illegalData() Dataset {
	return Dataset{name: "illegal", input: "\"test\n", tokens: []Token{
		{TokenIllegal, "test", 0, 0},
	}}
}
