package parser

import (
	"errors"

	"github.com/nurtai325/truth-table/internal/scanner"
)

var (
	ErrSyntax       = errors.New("syntax error")
	ErrIllegalToken = errors.New("illegal token")
)

func Parse(tokens []scanner.Token) (*Ast, error) {
	if len(tokens) == 0 {
		return nil, nil
	}
	if len(tokens) == 1 {
		return &Ast{Tok: Token{tok: tokens[0]}}, nil
	}

	i := operIndex(tokens)
	ast := new(Ast)
	ast.Tok = Token{tok: tokens[i]}

	l, err := Parse(tokens[:i])
	if err != nil {
		return nil, err
	}
	r, err := Parse(tokens[i+1:])
	if err != nil {
		return nil, err
	}
	ast.Left = l
	ast.Right = r

	return ast, nil
}

func operIndex(tokens []scanner.Token) int {
	for i := 0; i < len(tokens); i++ {
		if scanner.IsOperator(tokens[i]) {
			return i
		}
	}
	return -1
}
