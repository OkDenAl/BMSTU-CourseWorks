package main

import "slices"

var parsingTable = map[string][]string{
	"INNER_NEW_LINES CODE":         {},
	"RULE_DEF RULES_MARKER":        {},
	"INNER_RULES REG":              {"RULE", "INNER_RULES"},
	"NEW_LINES NL":                 {"NL", "INNER_NEW_LINES"},
	"INNER_NEW_LINES CLOSE_CODE":   {},
	"RULE REG":                     {"REG", "RULE_DEF"},
	"RULE_DEF REG":                 {},
	"INNER_RULES NL":               {"NL", "INNER_RULES"},
	"RULE_DEF OPEN_CODE":           {"OPEN_CODE", "INNER_NEW_LINES", "CODE", "INNER_NEW_LINES", "CLOSE_CODE"},
	"DEF OPEN_DEF":                 {"OPEN_DEF", "NEW_LINES", "CODE", "NEW_LINES", "CLOSE_DEF"},
	"INNER_NEW_LINES CLOSE_DEF":    {},
	"INNER_NEW_LINES NL":           {"NL", "INNER_NEW_LINES"},
	"INNER_NEW_LINES REG":          {},
	"INNER_NEW_LINES RULES_MARKER": {},
	"RULE_DEF NL":                  {},
	"PROG OPEN_DEF":                {"DEF", "NEW_LINES", "RULES", "NEW_LINES", "CODE"},
	"RULES RULES_MARKER":           {"RULES_MARKER", "NEW_LINES", "RULE", "INNER_RULES", "RULES_MARKER"},
	"INNER_RULES RULES_MARKER":     {},
}

var grammarAxiom = "PROG"

var nonTerms = []string{
	"PROG",
	"RULES",
	"INNER_RULES",
	"RULE",
	"RULE_DEF",
	"DEF",
	"NEW_LINES",
	"INNER_NEW_LINES",
}

func isTerminal(s string) bool {
	return !slices.Contains(nonTerms, s)
}
