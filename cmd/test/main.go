package main

import (
	"fmt"
	"log"

	"github.com/nurtai325/truth-table/internal/parser"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	exp := "!(!(a||b)||c||(a||c)&&(a||!(a||b)))"
	// exp := "(!a||b)->(b&&d)<=>b"
	// exp := "!(a||b)||c"
	// exp := "(a||b)||c"
	// exp := "a||b||c"
	// exp := "(a||b||c)"
	// exp := "a||b"
	// exp := "a || b || (!c || !d)"
	// exp := "!a || !b -> b && d <=> !b"
	fmt.Println(exp)
	ast, err := parser.Parse(&exp, true)
	if err != nil {
		log.Fatal(err)
	}
	ast.DfsWalk(func(ast *parser.Ast) {
		log.Println(ast)
	})
	// table, err := parser.Eval(ast)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(table)
}
