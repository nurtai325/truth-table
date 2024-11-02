package parser

import (
	"fmt"

	"github.com/nurtai325/truth-table/internal/scanner"
)

type Ast struct {
	Left    *Ast
	Right   *Ast
	Tok     scanner.Token
	Lit     string
	Negated bool
}

func (ast *Ast) Walk(f func(ast *Ast)) {
	if ast == nil {
		return
	}
	ast.Left.Walk(f)
	f(ast)
	ast.Right.Walk(f)
}

func (a Ast) String() string {
	if scanner.IsOperator(a.Tok) {
		return fmt.Sprintf("%v", a.Tok)
	}
	return fmt.Sprintf("%v %s negated: %t", a.Tok, a.Lit, a.Negated)
}
