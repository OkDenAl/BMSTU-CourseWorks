package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	inputFile  = flag.String("i", "./examples/lab1.2/example.gl", "input file location")
	outputFile = flag.String("o", "./examples/lab1.2/golexgen/lexer.go", "output file location")
	printTree  = flag.Bool("t", false, "if true -> print tree")
)

func main() {

	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	compiler := NewCompiler()
	scn := NewScanner(reader, &compiler)

	t := scn.NextToken()
	var tokens []Token
	for t.Tag() != TagEOP {
		if t.Tag() != TagErr {
			tokens = append(tokens, t)
		}
		t = scn.NextToken()
	}
	tokens = append(tokens, t)

	compiler.OutputMessages()
	parser := New(tokens)

	parse, err := parser.Parse()
	if err != nil {
		panic(err.Error())
	}

	var automatas []*FiniteState
	for _, rule := range parse.rules {
		automatas = append(automatas, rule.expr.Compile())
	}

	fmt.Println([]rune("\t"))
	fmt.Println(automatas[0].FindMatchEndIndex("\n"))

	fmt.Println(parse)

	gen := parse.Process()

	generateFile("templates/lexer.tmpl", *outputFile, gen)
}
