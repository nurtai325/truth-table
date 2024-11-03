package main

import (
	"fmt"
	"log"

	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	// exp := "!(a || !b) -> (b && d) <=> !b"
	// exp := "(a||b)->(b&&d)<=>b"
	// exp := "(a||b)||c"
	// exp := "a||b||c"
	exp := "(a||b)"
	// exp := "a||b"
	// exp := "a || b || !c || !d"
	fmt.Println(exp)
	ast, err := parser.Parse(&exp)
	if err != nil {
		log.Fatal(err)
	}
	ast.DfsWalk(func(ast *parser.Ast) {
		fmt.Println(ast)
	})
	ast.Print()
}
