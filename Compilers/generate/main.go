package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode"
)

type Tag int

const (
	ErrTag Tag = iota
	IdentTag
	NumTag
	AssemblyCommandTag
)

var tagToString = map[Tag]string{
	IdentTag:           "IDENT",
	NumTag:             "NUM",
	AssemblyCommandTag: "ASSEMBLY",
	ErrTag:             "ERR",
}

var (
	EOPTag  = Tag(-1)
	SKIPTag = Tag(-2)
)

type Token struct {
	tag    Tag
	coords Fragment
	val    string
}

func NewToken(tag Tag, starting, following Position, val string) Token {
	return Token{tag: tag, coords: NewFragment(starting, following), val: val}
}

type Fragment struct {
	starting  Position
	following Position
}

func NewFragment(starting, following Position) Fragment {
	return Fragment{starting: starting, following: following}
}

func (f Fragment) String() string {
	return fmt.Sprintf("%s-%s", f.starting.String(), f.following.String())
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s: %s", tagToString[t.tag], t.coords, t.val)
}

type Position struct {
	line  int
	pos   int
	index int
	text  []rune
}

func NewPosition(text []rune) Position {
	return Position{text: text, line: 1, pos: 1}
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.line, p.pos)
}

func (p *Position) Cp() int {
	if p.index == len(p.text) {
		return -1
	}
	return int(p.text[p.index])
}

func (p *Position) IsWhiteSpace() bool {
	return p.Cp() == ' '
}

func (p *Position) IsLetter() bool {
	return unicode.IsLetter(rune(p.Cp()))
}

func (p *Position) IsNewLine() bool {
	return p.Cp() == '\n'
}

func (p *Position) Next() Position {
	if p.index < len(p.text) {
		if p.IsNewLine() {
			p.line++
			p.pos = 1
		} else {
			p.pos++
		}
		p.index++
	}

	return *p
}

type Scanner struct {
	program []rune
	curPos  Position
}

func NewScanner(program []rune) Scanner {
	return Scanner{program: program, curPos: NewPosition(program)}
}

var (
	reg0 = regexp.MustCompile(`^[\t\n ]+`)
	reg1 = regexp.MustCompile(`^[a-z]*`)
	reg2 = regexp.MustCompile(`^[0-9]*`)
	reg3 = regexp.MustCompile(`^(mov|eax)`)
)

func (s *Scanner) findTag(maxLeftReg string) Tag {
	switch maxLeftReg {
	case `^[\t\n ]+`:
		fmt.Println("skipped", s.curPos.String())
	case `^[a-z]*`:
		return IdentTag
	case `^[0-9]*`:
		return NumTag
	case `^(mov|eax)`:
		return AssemblyCommandTag

	}
	return SKIPTag
}

func (s *Scanner) NextToken() Token {
	for s.curPos.Cp() != -1 {
		start := s.curPos.index

		maxRightReg := ""
		maxRight := 0
		regexps := make([]*regexp.Regexp, 0, 4)
		regexps = append(regexps, reg0)
		regexps = append(regexps, reg1)
		regexps = append(regexps, reg2)
		regexps = append(regexps, reg3)

		for _, r := range regexps {
			res := r.FindStringIndex(string(s.program[s.curPos.index:]))
			if res != nil && res[1] > maxRight {
				maxRightReg = r.String()
				maxRight = res[1]
			}
		}
		pos := s.curPos
		for s.curPos.index != start+maxRight {
			s.curPos.Next()
		}

		if maxRight == 0 {
			pos := s.curPos
			if s.curPos.Cp() != -1 {
				s.curPos.Next()
			} else {
				break
			}
			return NewToken(
				ErrTag,
				pos,
				pos,
				fmt.Sprintf("unknown symbol \"%s\"", string(s.program[start])),
			)
		} else {
			tok := s.findTag(maxRightReg)
			if tok != SKIPTag {
				return NewToken(tok, pos, s.curPos, string(s.program[pos.index:s.curPos.index]))
			}
		}
	}

	return NewToken(EOPTag, s.curPos, s.curPos, "")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage must be: go run main.go <fileTag.txt>\n")
	}
	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	scn := NewScanner([]rune(string(content)))

	t := scn.NextToken()
	for t.tag != EOPTag {
		fmt.Println(t.String())
		t = scn.NextToken()
	}
}
