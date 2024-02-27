package internal

import (
	"unicode"
)

type tokenId int

const (
	tokenEOF tokenId = iota
	tokenNil
	tokenBool
	tokenFloat
	tokenInt
	tokenString
	tokenSymbol
	tokenLeftParen
	tokenRightParen
	tokenIllegal
)

type token struct {
	id    tokenId
	value string
	row   int
	col   int
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

func keywordId(s string) (tokenId, string) {
	switch {
	case s == "true" || s == "false":
		return tokenBool, s
	case s == "nil":
		return tokenNil, ""
	}

	return tokenIllegal, ""
}
