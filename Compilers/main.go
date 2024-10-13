package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var (
	inputFile  = flag.String("i", "./examples/example.gl", "input file location")
	outputFile = flag.String("o", "./generate/main.go", "output file location")
	printTree  = flag.Bool("t", false, "if true -> print tree")
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	compiler := NewCompiler()
	scn := NewScanner(reader, &compiler)

	compiler.OutputMessages()

	parser := NewParser()

	tree, err := parser.TopDownParse(&scn)
	if err != nil {
		log.Panic(err)
	}

	if *printTree {
		tree.Print("")
	}

	gen := Interpret(tree)

	generateFile("templates/lexer.tmpl", *outputFile, gen)
}
