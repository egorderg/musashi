package internal

import (
	"bufio"
	"io"
	"unicode"
)

type lexer struct {
	row    int
	col    int
	colls  int
	reader *bufio.Reader
}

func newLexer(r io.Reader) *lexer {
	return &lexer{reader: bufio.NewReader(r)}
}

func (l *lexer) next() token {
	for {
		r, eof := l.readRune()
		if eof {
			return l.newToken(tokenEOF, "", 1)
		}

		if unicode.IsSpace(r) {
			if isNewLine(r) {
				l.resetPos()
			}
			continue
		}

		if isComment(r) {
			l.readTill(func(i int, r rune) bool {
				if isNewLine(r) {
					return false
				}

				return true
			})
			continue
		}

		return l.lex(r)
	}
}

func (l *lexer) lex(r rune) token {
	switch {
	case isLeftParen(r):
		return l.newToken(tokenLeftParen, "", 1)
	case isRightParen(r):
		return l.newToken(tokenRightParen, "", 1)
	case isNumber(r):
		l.unreadRune()
		return l.lexNumber()
	case isStringQuote(r):
		return l.lexString()
	default:
		l.unreadRune()
		return l.lexSymbol()
	}
}

func (l *lexer) lexNumber() token {
	isFloat := false
	v := l.readTill(func(i int, r rune) bool {
		if isFloatPoint(r) && !isFloat {
			isFloat = true
			return true
		}

		return isNumber(r)
	})

	if isFloat {
		return l.newToken(tokenFloat, string(v), len(v))
	} else {
		return l.newToken(tokenInt, string(v), len(v))
	}
}

func (l *lexer) lexString() token {
	isInvalid := false
	v := l.readTill(func(i int, r rune) bool {
		if isNewLine(r) {
			isInvalid = true
			return false
		}

		return !isStringQuote(r)
	})

	s := string(v)

	if isInvalid {
		return l.newToken(tokenIllegal, s, len(s)+1)
	} else {
		l.readRune()
	}

	return l.newToken(tokenString, s, len(s)+2)
}

func (l *lexer) lexSymbol() token {
	v := l.readTill(func(_ int, r rune) bool {
		return !isSpace(r) && !isStringQuote(r) && !isNumber(r) && !isComment(r)
	})

	s := string(v)
	if t, tval := keywordId(s); t != tokenIllegal {
		return l.newToken(t, tval, len(s))
	}

	return l.newToken(tokenSymbol, s, len(s))
}

func (l *lexer) readTill(cb func(i int, r rune) bool) []rune {
	v := make([]rune, 0, 8)
	i := 0

	for {
		r, eof := l.readRune()
		if eof {
			break
		}

		if cb(i, r) {
			v = append(v, r)
		} else {
			l.unreadRune()
			break
		}

		i++
	}

	return v
}

func (l *lexer) readRune() (rune, bool) {
	for {
		r, s, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return '\x00', true
			}

			panic(err)
		}

		l.col += s
		l.colls = s
		return r, false
	}
}

func (l *lexer) unreadRune() {
	l.col -= l.colls
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
}

func (l *lexer) resetPos() {
	l.col = 0
	l.row++
}

func (l *lexer) newToken(id tokenId, v string, length int) token {
	return token{id: id, value: v, col: l.col - length, row: l.row}
}
