package main

import (
	"fmt"
	"log"

	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	exp := "!(a || !b) -> (b && d) <=> !b"
	// exp := "(a||b)->(b&&d)<=>b"
	// exp := "(a||b)||c"
	// exp := "(a||b)"
	// exp := "a||b||c"
	ast, _, err := parser.Parse(&exp)
	if err != nil {
		log.Fatal(err)
	}
	ast.Walk(func(ast *parser.Ast) {
		fmt.Println(ast)
	})
}
