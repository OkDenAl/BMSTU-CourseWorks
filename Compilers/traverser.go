package main

import (
	"fmt"
	"log"
)

type Regexp struct {
	RegexpName string
	RegexpVal  string
	Action     string
}

type GeneratorInfo struct {
	DefCode    string
	UserCode   string
	Regexps    []Regexp
	RegexpsLen int
}

//$AXIOM PROG
//$NTERM RULES INNER_RULES RULE RULE_DEF DEF NEW_LINES INNER_NEW_LINES
//$TERM "OPEN_DEF" "CLOSE_DEF" "RULES_MARKER" "OPEN_CODE" "CLOSE_CODE" "CODE" "NL" "REG"
//
//$RULE PROG = DEF NEW_LINES RULES NEW_LINES "CODE"
//$RULE DEF = "OPEN_DEF" NEW_LINES "CODE" NEW_LINES "CLOSE_DEF"
//$RULE RULES = "RULES_MARKER" NEW_LINES RULE INNER_RULES "RULES_MARKER"
//$RULE INNER_RULES = "NL" INNER_RULES
//RULE INNER_RULES
//$EPS
//$RULE NEW_LINES = "NL" INNER_NEW_LINES
//$RULE INNER_NEW_LINES = "NL" INNER_NEW_LINES
//$EPS
//$RULE RULE = "REG" RULE_DEF
//$RULE RULE_DEF = "OPEN_CODE" INNER_NEW_LINES "CODE" INNER_NEW_LINES "CLOSE_CODE"
//$EPS

func Interpret(tree NodePrinter) GeneratorInfo {
	info := GeneratorInfo{
		Regexps: make([]Regexp, 0),
	}

	rulesCount := 0
	var traverse func(treeNode NodePrinter)
	traverse = func(treeNode NodePrinter) {
		if node, ok := treeNode.(*InnerNode); ok {
			switch node.nterm {
			case "PROG":
				if len(node.children) == 5 {
					traverse(node.children[0])
					traverse(node.children[2])
					if val, ok := node.children[4].(*Leaf); ok {
						info.UserCode = val.tok.val
					} else {
						log.Fatal("Невалидное дерево, ожидался лист")
					}
				} else {
					log.Fatal("Невалидная длина PROG")
				}
			case "DEF":
				if len(node.children) == 5 {
					if val, ok := node.children[2].(*Leaf); ok {
						info.DefCode = val.tok.val
					} else {
						log.Fatal("Невалидное дерево, ожидался лист")
					}
				} else {
					log.Fatal("Невалидная длина DEF")
				}
			case "RULES":
				if len(node.children) == 5 {
					traverse(node.children[2])
					traverse(node.children[3])
				} else {
					log.Fatal("Невалидная длина RULES")
				}
			case "INNER_RULES":
				if len(node.children) == 2 {
					if _, ok = node.children[0].(*InnerNode); ok {
						traverse(node.children[0])
					}
					traverse(node.children[1])
				} else if len(node.children) != 0 {
					log.Fatal("Невалидная длина INNER_RULES ")
				}
			case "NEW_LINES":
				if len(node.children) != 2 {
					log.Fatal("Невалидная длина NEW_LINES ")
				}
			case "INNER_NEW_LINES":
				if len(node.children) != 2 && len(node.children) != 0 {
					log.Fatal("Невалидная длина INNER_NEW_LINES ")
				}
			case "RULE":
				if len(node.children) == 2 {
					if val, ok := node.children[0].(*Leaf); ok {
						info.Regexps = append(info.Regexps, Regexp{
							RegexpVal:  "^" + val.tok.val,
							RegexpName: fmt.Sprintf("reg%d", rulesCount),
						})
						rulesCount++
					} else {
						log.Fatal("Невалидное дерево, ожидался лист")
					}
					traverse(node.children[1])
				} else {
					log.Fatal("Невалидная длина RULE")
				}
			case "RULE_DEF":
				if len(node.children) == 5 {
					if val, ok := node.children[2].(*Leaf); ok {
						info.Regexps[rulesCount-1].Action = val.tok.val
					} else {
						log.Fatal("Невалидное дерево, ожидался лист")
					}
				} else if len(node.children) != 0 {
					log.Fatal("Невалидная длина RULE_DEF")
				}
			default:
				log.Fatal("Неизвестный нетерминал", node, len(node.children))
			}
		}
	}
	traverse(tree)
	info.RegexpsLen = len(info.Regexps)

	return info
}
