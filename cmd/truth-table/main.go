package main

import (
	"flag"
	"log"

	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	log.SetFlags(0)

	exp := flag.String("e", "", "specify the `logic expression` to be parsed.")
	printAst := flag.Bool("ast", false, "print the `ast` of the expression")
	subExpressions := flag.Bool("sub", false, "include `sub expressions` to the truth table too")
	flag.Parse()

	ast, err := parser.Parse(exp)
	if err != nil {
		log.Fatal(err)
	}
	table, err := ast.Eval(*subExpressions)
	if err != nil {
		log.Fatal(err)
	}
	table.Print()
	if *printAst {
		ast.Print()
	}
}
