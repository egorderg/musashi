package lexer

import (
	"unicode"
)

const (
	TokenEOF = iota
	TokenNil
	TokenBool
	TokenNumber
	TokenString
	TokenSymbol
	TokenLeftParen
	TokenRightParen
	TokenIllegal
)

type Token struct {
	Type   int
	Value  any
	Row    int
	Column int
}

func isNewLine(r rune) bool {
	return r == '\n'
}

func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func isStringQuote(r rune) bool {
	return r == '"'
}

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}

func isFloatPoint(r rune) bool {
	return r == '.'
}

func isComment(r rune) bool {
	return r == '#'
}

func isLeftParen(r rune) bool {
	return r == '('
}

func isRightParen(r rune) bool {
	return r == ')'
}

func keywordType(s string) (int, any) {
	switch {
	case s == "true" || s == "false":
		return TokenBool, s
	case s == "nil":
		return TokenNil, nil
	}

	return TokenIllegal, nil
}
