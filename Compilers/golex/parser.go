package main

import "fmt"

type Parser struct {
	cursor int
	tokens []Token
}

func New(tokens []Token) *Parser {
	return &Parser{
		cursor: 0,
		tokens: tokens,
	}
}

// Parse a list of tokens into an executable
func (p *Parser) Parse() (*Rules, error) {
	return p.rules()
}

// Rules           ::= "%%" (NewLine)+ (Rule)+ "%%"
func (p *Parser) rules() (*Rules, error) {
	p.mustExpectTags(TagRulesMarker)
	p.mustExpectTags(TagNL)
	for p.tokens[p.cursor].tag == TagNL {
		p.mustExpectTags(TagNL)
	}

	var rules []*Rule
	for p.tokens[p.cursor].tag == TagRegularMarker {
		rule, err := p.rule()
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	p.mustExpectTags(TagRulesMarker)

	return &Rules{rules: rules}, nil
}

// Rule            ::= "/" RegExpr "/" TagRuleName NL
func (p *Parser) rule() (*Rule, error) {
	p.mustExpectTags(TagRegularMarker)
	expr, ok := p.regExpr()
	if !ok {
		return nil, fmt.Errorf("parse error: failed to parse regular expr starts with %s", p.tokens[p.cursor])
	}
	p.mustExpectTags(TagRegularMarker)
	token := p.mustExpectTags(TagRuleName)
	p.mustExpectTags(TagNL)
	for p.tokens[p.cursor].tag == TagNL {
		p.mustExpectTags(TagNL)
	}

	return &Rule{expr: expr, name: &token}, nil
}

func (p *Parser) regExpr() (*RegExpr, bool) {
	reset := p.reset()
	union, ok := p.union()
	if ok {
		return &RegExpr{union: union}, true
	}

	reset()
	simple, ok := p.simpleExpr()
	if ok {
		return &RegExpr{simple: simple}, true
	}

	return nil, false
}

func (p *Parser) union() (*Union, bool) {
	simple, ok := p.simpleExpr()
	if !ok {
		return nil, false
	}

	if _, ok = p.expectTags(TagPipe); !ok {
		return nil, false
	}

	regex, ok := p.regExpr()

	if !ok {
		return nil, false
	}

	return &Union{regex, simple}, true
}

func (p *Parser) simpleExpr() (*SimpleExpr, bool) {
	concatenation, ok := p.concatenation()
	if ok {
		return &SimpleExpr{concatenation: concatenation}, true
	}

	basic, ok := p.basicExpr()
	if ok {
		return &SimpleExpr{basic: basic}, true
	}

	return nil, false
}

func (p *Parser) concatenation() (*Concatenation, bool) {
	reset := p.reset()

	basic, ok := p.basicExpr()

	if !ok {
		reset()
		return nil, false
	}

	simple, ok := p.simpleExpr()

	if !ok {
		reset()
		return nil, false
	}

	return &Concatenation{simple, basic}, true

}

func (p *Parser) basicExpr() (*BasicExpr, bool) {
	base, ok := p.element()
	if !ok {
		return nil, false
	}

	token, ok := p.expectTags(TagStar, TagPlus, TagQuestion)
	if ok {
		return &BasicExpr{element: base, op: &token}, true
	}

	return &BasicExpr{element: base}, true
}

func (p *Parser) element() (*Element, bool) {
	group, ok := p.group()
	if ok {
		return &Element{group: group}, true
	}

	set, ok := p.set()
	if ok {
		return &Element{set: set}, true
	}

	character, ok := p.character()
	if ok {
		return &Element{character: character}, true
	}

	escape, ok := p.escape()
	if ok {
		return &Element{escape: escape}, true
	}

	return nil, false
}

func (p *Parser) group() (*Group, bool) {
	if _, ok := p.expectTags(TagOpenParen); !ok {
		return nil, false
	}

	regex, ok := p.regExpr()

	if !ok {
		return nil, false
	}

	if _, ok = p.expectTags(TagCloseParen); !ok {
		return nil, false
	}

	return &Group{regex}, true
}

func (p *Parser) escape() (*Escape, bool) {
	if _, ok := p.expectTags(TagEscape); !ok {
		return nil, false
	}

	base, ok := p.token()
	if !ok {
		return nil, false
	}

	return &Escape{base}, true

}

func (p *Parser) set() (*Set, bool) {
	reset := p.reset()
	if _, ok := p.expectTags(TagOpenBracket); !ok {
		return nil, false
	}

	var set *Set
	if _, ok := p.expectTags(TagCaret); !ok {
		positive, ok := p.setItems()
		if ok {
			set = &Set{positive: positive}
		}
	} else {
		negative, ok := p.setItems()
		if ok {
			set = &Set{negative: negative}
		}
	}

	if set != nil {
		p.mustExpectTags(TagCloseBracket)
		return set, true
	}

	reset()
	return nil, false

}

func (p *Parser) setItems() (*SetItems, bool) {
	item, ok := p.setItem()

	if !ok {
		return nil, false
	}

	items, ok := p.setItems()

	return &SetItems{item: item, items: items}, true

}

func (p *Parser) setItem() (*SetItem, bool) {
	reset := p.reset()
	rnge, ok := p.rangeExpr()
	if ok {
		return &SetItem{rnge: rnge}, true
	}

	reset()
	escape, ok := p.escape()
	if ok {
		return &SetItem{escape: escape}, true
	}

	reset()
	token, ok := p.character(withTag(TagDash))
	if ok {
		return &SetItem{character: token}, true
	}

	reset()
	return nil, false

}

func (p *Parser) rangeExpr() (*Range, bool) {
	start, ok := p.token()

	if !ok {
		return nil, false
	}

	if _, ok = p.expectTags(TagDash); !ok {
		return nil, false
	}

	end, ok := p.character()

	if !ok {
		return nil, false
	}

	return &Range{start, end}, true

}

type characterOpts func(*Token) (*Character, bool)

func withTag(tag DomainTag) characterOpts {
	return func(token *Token) (*Character, bool) {
		if token.Tag() == tag {
			return &Character{token}, true
		}

		return nil, false
	}
}

func (p *Parser) character(opts ...characterOpts) (*Character, bool) {
	reset := p.reset()

	base, ok := p.token()

	if !ok {
		return nil, false
	}

	for _, opt := range opts {
		if res, ok := opt(base); ok {
			return res, true
		}
	}

	if res, ok := withTag(TagCharacter)(base); ok {
		return res, true
	}
	if res, ok := withTag(TagAnyCharacter)(base); ok {
		return res, true
	}

	reset()
	return nil, false
}

func (p *Parser) token() (*Token, bool) {
	token, ok := p.nextToken()
	if !ok {
		return nil, false
	}

	return &token, true
}

func (p *Parser) nextToken() (Token, bool) {
	if p.cursor == len(p.tokens) {
		if len(p.tokens) == 0 {
			return Token{tag: TagCharacter, val: " "}, false
		}
		return p.tokens[p.cursor-1], false
	}

	token := p.tokens[p.cursor]
	p.cursor++

	return token, true
}

func (p *Parser) mustExpectTags(tags ...DomainTag) Token {
	for _, tag := range tags {
		if p.tokens[p.cursor].tag == tag {
			token := p.tokens[p.cursor]
			p.cursor = p.cursor + 1
			if p.tokens[p.cursor].tag == TagErr {
				panic(fmt.Sprintf("parse error: unexpected token"))
			}
			return token
		}
	}

	panic(fmt.Sprintf("parse error: expected %s, but got %s", tags, p.tokens[p.cursor]))
}

func (p *Parser) expectTags(tags ...DomainTag) (Token, bool) {
	reset := p.reset()
	token, ok := p.nextToken()
	if !ok {
		reset()
		return Token{}, false
	}

	for _, tag := range tags {
		if token.Tag() == tag {
			return token, true
		}
	}

	reset()
	return Token{}, false
}

func (p *Parser) reset() func() {
	cursor := p.cursor
	return func() {
		p.cursor = cursor
	}
}
