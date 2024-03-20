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

func (p *parser) parse() (AstProgram, bool) {
	forms := make([]AstForm, 0)

	for {
		item, _ := p.parseItem()
		if form, ok := item.(AstForm); ok {
			forms = append(forms, form)
		} else {
			// TODO: error
			break
		}
	}

	if len(p.errors) > 0 {
		return AstProgram{}, false
	}

	return AstProgram{forms: forms}, true
}

func (p *parser) parseItem() (AstItem, token) {
	t := p.lexer.next()
	if t.id == tokenEOF {
		return nil, t
	}

	switch t.id {
	case tokenNil:
		return AstDatum{id: datumNil, value: t.value}, t
	case tokenInt:
		return AstDatum{id: datumInt, value: t.value}, t
	case tokenFloat:
		return AstDatum{id: datumFloat, value: t.value}, t
	case tokenBool:
		return AstDatum{id: datumBool, value: t.value}, t
	case tokenString:
		return AstDatum{id: datumString, value: t.value}, t
	case tokenSymbol:
		return AstSymbol{name: t.value}, t
	case tokenLeftParen:
		if f, ok := p.parseForm(); ok {
			return f, t
		}
		return nil, t
	default:
		return nil, t
	}
}

func (p *parser) parseForm() (AstItem, bool) {
	items := make([]AstItem, 0)

	for {
		item, tok := p.parseItem()
		if tok.id == tokenEOF {
			// TODO: error
			return nil, false
		}

		if tok.id == tokenRightParen {
			break
		}

		items = append(items, item)
	}

	return AstForm{items: items}, true
}

func (p *parser) nextToken(want tokenId) (token, bool) {
	tok := p.lexer.next()

	if tok.id != want {
		p.errors = append(p.errors, ParseError{got: tok, want: want})
		return tok, tok.id == tokenEOF
	}

	return tok, false
}
