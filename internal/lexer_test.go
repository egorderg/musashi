package internal

import (
	"bufio"
	"strings"
	"testing"
)

type Dataset struct {
	name   string
	input  string
	tokens []token
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

			lexer := newLexer(bufio.NewReader(strings.NewReader(ds.input)))
			tokens := make([]token, 0, 10)

			for {
				tok := lexer.next()
				if tok.id == tokenEOF {
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
	return Dataset{
		name:  "comments",
		input: "true#k平仮 2.234 dk jfk \n1.23#test",
		tokens: []token{
			{tokenBool, "true", 0, 0},
			{tokenFloat, "1.23", 1, 0},
		},
	}
}

func keywordData() Dataset {
	return Dataset{
		name:  "keywords",
		input: "true false nil",
		tokens: []token{
			{tokenBool, "true", 0, 0},
			{tokenBool, "false", 0, 5},
			{tokenNil, "", 0, 11},
		},
	}
}

func symbolData() Dataset {
	return Dataset{
		name:  "symbols",
		input: "++\"str\"平032^t#test\n-a.bc :po\r\nset _ + * & | < > : % / -\n %-平仮-*k/",
		tokens: []token{
			{tokenSymbol, "++", 0, 0},
			{tokenString, "str", 0, 2},
			{tokenSymbol, "平", 0, 7},
			{tokenInt, "032", 0, 10},
			{tokenSymbol, "^t", 0, 13},
			{tokenSymbol, "-a.bc", 1, 0},
			{tokenSymbol, ":po", 1, 6},
			{tokenSymbol, "set", 2, 0},
			{tokenSymbol, "_", 2, 4},
			{tokenSymbol, "+", 2, 6},
			{tokenSymbol, "*", 2, 8},
			{tokenSymbol, "&", 2, 10},
			{tokenSymbol, "|", 2, 12},
			{tokenSymbol, "<", 2, 14},
			{tokenSymbol, ">", 2, 16},
			{tokenSymbol, ":", 2, 18},
			{tokenSymbol, "%", 2, 20},
			{tokenSymbol, "/", 2, 22},
			{tokenSymbol, "-", 2, 24},
			{tokenSymbol, "%-平仮-*k/", 3, 1},
		},
	}
}

func numberData() Dataset {
	return Dataset{
		name:  "numbers",
		input: "05 7 -13\n47 -1.34\r\n27.234-b",
		tokens: []token{
			{tokenInt, "05", 0, 0},
			{tokenInt, "7", 0, 3},
			{tokenSymbol, "-", 0, 5},
			{tokenInt, "13", 0, 6},
			{tokenInt, "47", 1, 0},
			{tokenSymbol, "-", 1, 3},
			{tokenFloat, "1.34", 1, 4},
			{tokenFloat, "27.234", 2, 0},
			{tokenSymbol, "-b", 2, 6},
		},
	}
}

func stringData() Dataset {
	return Dataset{
		name:  "strings",
		input: "\"\" \"平仮名\"\"value\"\n \"long string\"",
		tokens: []token{
			{tokenString, "", 0, 0},
			{tokenString, "平仮名", 0, 3},
			{tokenString, "value", 0, 14},
			{tokenString, "long string", 1, 1},
		},
	}
}

func complexData() Dataset {
	return Dataset{
		name:  "complex",
		input: "#fn\r\n(+ 2 \n(!pos :pos-平 -2.7 \"hello 平\" 5\n))",
		tokens: []token{
			{tokenLeftParen, "", 1, 0},
			{tokenSymbol, "+", 1, 1},
			{tokenInt, "2", 1, 3},
			{tokenLeftParen, "", 2, 0},
			{tokenSymbol, "!pos", 2, 1},
			{tokenSymbol, ":pos-平", 2, 6},
			{tokenSymbol, "-", 2, 15},
			{tokenFloat, "2.7", 2, 16},
			{tokenString, "hello 平", 2, 20},
			{tokenInt, "5", 2, 32},
			{tokenRightParen, "", 3, 0},
			{tokenRightParen, "", 3, 1},
		},
	}
}

func concatData() Dataset {
	return Dataset{
		name:  "concat",
		input: "\"test 555\".34.23.\"\"平35-3;4\n4/!pos7true9\"test平\"\n#test\n\"test2\"",
		tokens: []token{
			{tokenString, "test 555", 0, 0},
			{tokenSymbol, ".", 0, 10},
			{tokenFloat, "34.23", 0, 11},
			{tokenSymbol, ".", 0, 16},
			{tokenString, "", 0, 17},
			{tokenSymbol, "平", 0, 19},
			{tokenInt, "35", 0, 22},
			{tokenSymbol, "-", 0, 24},
			{tokenInt, "3", 0, 25},
			{tokenSymbol, ";", 0, 26},
			{tokenInt, "4", 0, 27},
			{tokenInt, "4", 1, 0},
			{tokenSymbol, "/!pos", 1, 1},
			{tokenInt, "7", 1, 6},
			{tokenBool, "true", 1, 7},
			{tokenInt, "9", 1, 11},
			{tokenString, "test平", 1, 12},
			{tokenString, "test2", 3, 0},
		},
	}
}

func illegalData() Dataset {
	return Dataset{
		name:  "illegal",
		input: "\"test\n",
		tokens: []token{
			{tokenIllegal, "test", 0, 0},
		},
	}
}
