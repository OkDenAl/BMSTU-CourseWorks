package main

import (
	"bufio"
	"fmt"
)

type comment struct {
	Fragment
	value string
}

func newComment(starting, following Position, value string) comment {
	return comment{NewFragment(starting, following), value}
}

func (c comment) String() string {
	return fmt.Sprintf("COMMENT %s-%s: %s", c.starting.String(), c.following.String(), c.value)
}

type Scanner struct {
	programReader *bufio.Reader
	compiler      *Compiler
	curPos        Position
	comments      []comment
}

func NewScanner(programFile *bufio.Reader, compiler *Compiler) Scanner {
	return Scanner{programReader: programFile, compiler: compiler, curPos: NewPosition(programFile)}
}

func (s *Scanner) printComments() {
	for _, comm := range s.comments {
		fmt.Println(comm)
	}
}

func (s *Scanner) NextToken() Token {
	for s.curPos.Cp() != -1 {
		for s.curPos.IsWhiteSpace() {
			s.curPos.Next()
		}
		start := s.curPos
		curWord := ""

		switch s.curPos.Cp() {
		case '\n':
			s.curPos.Next()
			return NewToken(TagNL, start, start, "NEW_LINE")
		case '%':
			curWord += string(rune(s.curPos.Cp()))
			s.curPos.Next()

			if s.curPos.Cp() == -1 || (s.curPos.GetSymbol() != '{' &&
				s.curPos.GetSymbol() != '%' && s.curPos.GetSymbol() != '}') {

				s.compiler.AddMessage(true, start, "invalid syntax")
				s.curPos.SkipErrors()

				return NewToken(TagErr, s.curPos, s.curPos, "")
			}

			curWord += string(rune(s.curPos.Cp()))
			pos := s.curPos
			s.curPos.Next()

			var token Token
			switch curWord {
			case "%%":
				token = NewToken(TagRulesMarker, start, pos, curWord)
			case "%{":
				token = NewToken(TagOpenDef, start, pos, curWord)
			case "%}":
				token = NewToken(TagCloseDef, start, pos, curWord)
			}

			return token
		case '{':
			s.curPos.Next()
			return NewToken(TagOpenCode, start, start, "{")
		case '}':
			s.curPos.Next()
			return NewToken(TagCloseCode, start, start, "}")
		case '[':
			var bracketBalance int
			for s.curPos.Cp() != -1 {
				s.curPos.Next()
				if s.curPos.GetSymbol() == ']' && bracketBalance == 0 {
					break
				} else if s.curPos.GetSymbol() == '[' {
					bracketBalance++
				} else if s.curPos.GetSymbol() == ']' {
					bracketBalance--
				}
				curWord += string(s.curPos.GetSymbol())
			}

			if s.curPos.Cp() == -1 {
				s.compiler.AddMessage(true, start, "invalid syntax")
				s.curPos.SkipErrors()

				return NewToken(TagErr, s.curPos, s.curPos, "")
			}
			
			pos := s.curPos
			s.curPos.Next()

			return NewToken(TagRegExp, start, pos, curWord)
		default:
			var pos Position
			var bracketBalance int
			for s.curPos.Cp() != -1 {
				curWord += string(s.curPos.GetSymbol())
				pos = s.curPos
				s.curPos.Next()
				if s.curPos.GetSymbol() == '%' {
					pos1 := s.curPos
					s.curPos.Next()
					if s.curPos.GetSymbol() == '}' {
						var curPos Position
						if curWord[len(curWord)-1] == '\n' {
							curWord = curWord[:len(curWord)-1]
							looked = append(looked, pos1, s.curPos)
							curPos = pos
						} else {
							curPos = pos1
							looked = append(looked, s.curPos)
						}
						s.curPos = curPos
						break
					}
				} else if s.curPos.GetSymbol() == '}' && bracketBalance == 0 {
					break
				} else if s.curPos.GetSymbol() == '{' {
					bracketBalance++
				} else if s.curPos.GetSymbol() == '}' {
					bracketBalance--
				}
			}

			return NewToken(TagCode, start, pos, curWord)
		}
	}

	return NewToken(TagEOP, s.curPos, s.curPos, "")
}
