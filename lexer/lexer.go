package lexer

import (
	"bufio"
	"io"
	"unicode"
)

type Lexer struct {
	row    int
	col    int
	colls  int
	reader *bufio.Reader
}

func NewLexer(r *bufio.Reader) *Lexer {
	return &Lexer{reader: r}
}

func (l *Lexer) Next() Token {
	for {
		r, eof := l.readRune()
		if eof {
			return l.newToken(TokenEOF, nil, 1)
		}

		if unicode.IsSpace(r) {
			if r == '\n' {
				l.resetPos()
			}
			continue
		}

		if r == comment {
			l.readTill(func(i int, r rune) bool {
				if r == '\n' {
					return false
				}

				return true
			})
			continue
		}

		return l.lex(r)
	}
}

func (l *Lexer) lex(r rune) Token {
	switch {
	case r == leftParen:
		return l.newToken(TokenLeftParen, nil, 1)
	case r == rightParen:
		return l.newToken(TokenRightParen, nil, 1)
	case r == terminator:
		return l.newToken(TokenTerminator, nil, 1)
	case isNumber(r):
		l.unreadRune()
		return l.lexNumber()
	case isSymbolStart(r):
		l.unreadRune()
		return l.lexSymbol()
	case r == stringQuote:
		return l.lexString()
	default:
		return l.newToken(TokenIllegal, string(r), 1)
	}
}

func (l *Lexer) lexSymbol() Token {
	v := l.readTill(func(_ int, r rune) bool {
		return isSymbol(r)
	})

	s := string(v)
	switch {
	case s == valueTrue || s == valueFalse:
		return l.newToken(TokenBool, s, len(s))
	case s == valueNil:
		return l.newToken(TokenNil, nil, len(s))
	}

	return l.newToken(TokenSymbol, s, len(s))
}

func (l *Lexer) lexNumber() Token {
	isFloat := false
	v := l.readTill(func(i int, r rune) bool {
		if r == floatPoint && !isFloat {
			isFloat = true
			return true
		}

		return isNumber(r)
	})

	return l.newToken(TokenNumber, string(v), len(v))
}

func (l *Lexer) lexString() Token {
	isInvalid := false
	v := l.readTill(func(i int, r rune) bool {
		if r == '\n' {
			isInvalid = true
			return false
		}

		return r != stringQuote
	})

	s := string(v)

	if isInvalid {
		return l.newToken(TokenIllegal, s, len(s)+1)
	} else {
		l.readRune()
	}

	return l.newToken(TokenString, s, len(s)+2)
}

func (l *Lexer) readTill(cb func(i int, r rune) bool) []rune {
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

func (l *Lexer) readRune() (rune, bool) {
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

func (l *Lexer) unreadRune() {
	l.col -= l.colls
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
}

func (l *Lexer) resetPos() {
	l.col = 0
	l.row++
}

func (l *Lexer) newToken(t int, v any, length int) Token {
	return Token{Type: t, Value: v, Column: l.col - length, Row: l.row}
}
