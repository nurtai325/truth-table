package main

import (
	"fmt"

	"github.com/nurtai325/truth-table/internal/parser"
	"github.com/nurtai325/truth-table/internal/scanner"
)

func main() {
	var s scanner.Scanner
	s.SetSrc([]byte("a||b||c"))

	var tokens []scanner.Token
	for {
		tok, _, _ := s.Scan()
		if tok == scanner.EOF {
			break
		}
		tokens = append(tokens, tok)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(ast.Left.Tok)
	fmt.Println(ast.Tok)
	fmt.Println(ast.Right.Left.Tok)
	fmt.Println(ast.Right.Tok)
	fmt.Println(ast.Right.Right.Tok)
}
