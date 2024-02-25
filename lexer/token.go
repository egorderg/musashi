package lexer

import (
	"strings"
	"unicode"
)

const (
	terminator  = '.'
	leftParen   = '('
	rightParen  = ')'
	stringQuote = '"'
	floatPoint  = '.'
	comment     = '#'
	valueTrue   = "true"
	valueFalse  = "false"
	valueNil    = "nil"
	symbols     = "_-+*/:%!&|<>="
)

const (
	TokenEOF = iota
	TokenNil
	TokenBool
	TokenNumber
	TokenString
	TokenSymbol
	TokenTerminator
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

func isSymbol(r rune) bool {
	return isSymbolStart(r)
}

func isSymbolStart(r rune) bool {
	return unicode.IsLetter(r) || strings.ContainsRune(symbols, r)
}

func isNumber(r rune) bool {
	return unicode.IsDigit(r)
}
