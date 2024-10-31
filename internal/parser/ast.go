package parser

import "github.com/nurtai325/truth-table/internal/scanner"

type Token struct {
	tok scanner.Token
	lit string
}

type Ast struct {
	Left *Ast
	Right *Ast
	operFunc operFunc
	Tok Token
}

func (ast *Ast) DfsWalk(f func(ast *Ast)) {
	f(ast)
	if ast.Left != nil {
		ast.Left.DfsWalk(f)
	}
	if ast.Right != nil {
		ast.Right.DfsWalk(f)
	}
}

type operFunc func(a, b *bool) *bool
