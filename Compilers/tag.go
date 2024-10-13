package main

type DomainTag int

const (
	TagCode = iota
	TagOpenDef
	TagCloseDef
	TagRulesMarker
	TagOpenCode
	TagCloseCode
	TagErr
	TagRegExp
	TagNL
	TagEOP
)

var tagToString = map[DomainTag]string{
	TagCode:        "CODE",
	TagOpenDef:     "OPEN_DEF",
	TagCloseDef:    "CLOSE_DEF",
	TagRulesMarker: "RULES_MARKER",
	TagOpenCode:    "OPEN_CODE",
	TagCloseCode:   "CLOSE_CODE",
	TagRegExp:      "REG",
	TagNL:          "NL",
	TagEOP:         "EOP",
	TagErr:         "ERR",
}
