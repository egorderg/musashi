package internal

import (
	"io"
)

type parser struct {
	lexer  *lexer
	errors []ParseError
}

func newParser(r io.Reader) *parser {
	return &parser{
		lexer:  newLexer(r),
		errors: make([]ParseError, 0),
	}
}

func (p *parser) parse() (*AstProgram, bool) {
	forms := make([]*AstForm, 0)

	for {
		if form, eof := p.parseForm(); !eof {
			forms = append(forms, form)
		} else {
			break
		}
	}

	if len(p.errors) > 0 {
		return nil, false
	}

	return &AstProgram{forms: forms}, true
}

func (p *parser) parseForm() (*AstForm, bool) {
	items := make([]AstItem, 0)
	if _, eof := p.nextToken(tokenLeftParen); eof {
		return nil, true
	}

	for {
		tok := p.lexer.next()
		if tok.id == tokenEOF {
			return nil, true
		}

		if tok.id == tokenRightParen {
			break
		}

		items = append(items, p.parseItem(tok))
	}

	return &AstForm{items: items}, false
}

func (p *parser) parseItem(t token) AstItem {
	switch t.id {
	case tokenNil:
		return AstDatum{id: datumNil, value: t.value}
	case tokenInt:
		return AstDatum{id: datumInt, value: t.value}
	case tokenFloat:
		return AstDatum{id: datumFloat, value: t.value}
	case tokenBool:
		return AstDatum{id: datumBool, value: t.value}
	case tokenString:
		return AstDatum{id: datumString, value: t.value}
	case tokenSymbol:
		return AstSymbol{name: t.value}
	default:
		return nil
	}
}

func (p *parser) nextToken(want tokenId) (token, bool) {
	tok := p.lexer.next()

	if tok.id != want {
		p.errors = append(p.errors, ParseError{got: tok, want: want})
		return tok, tok.id == tokenEOF
	}

	return tok, false
}
