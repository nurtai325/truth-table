package parser

import (
	"errors"
	"fmt"

	"github.com/nurtai325/truth-table/internal/scanner"
)

var (
	ErrSyntax       = errors.New("syntax error")
	ErrIllegalToken = errors.New("illegal token")
)

type parser struct {
	tokens []scanner.Token
	vars   []string
	opers  []int
	cursor int
}

func Parse(exp *string) (ast *Ast, err error) {
	var s scanner.Scanner
	src := []byte(*exp)
	s.SetSrc(src)

	var tokens []scanner.Token
	var variables []string
	var operators []int
Loop:
	for i := 0; true; i++ {
		token, lit, err := s.Scan()
		if err != nil {
			return nil, err
		}

		switch token {
		case scanner.EOF:
			break Loop
		case scanner.ILLEGAL:
			return nil, ErrIllegalToken
		case scanner.NOT:
			break
		case scanner.LPAREN:
			break
		case scanner.RPAREN:
			break
		default:
			if scanner.IsOperator(token) {
				operators = append(operators, i)
			}
		}

		tokens = append(tokens, token)
		if token == scanner.VAR {
			variables = append(variables, lit)
		}
	}

	// TODO: uncomment for proper error handling
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		ast = nil
	// 		err = e.(error)
	// 	}
	// }()

	p := new(parser)
	p.init(tokens, operators, variables)
	ast = p.parse()
	return
}

func (p *parser) init(tokens []scanner.Token, opers []int, vars []string) {
	p.tokens = tokens
	p.opers = opers
	p.vars = vars
	p.cursor = 0
}

func (p *parser) parse() *Ast {
	fmt.Println(p.tokens, p.opers, p.vars)
	switch {
	case len(p.tokens) == 1:
		return p.parseVar()
	case len(p.tokens) == 2 && p.tokens[0] == scanner.NOT:
		p.tokens = p.tokens[1:]
		ast := p.parseVar()
		ast.negated = true
		return ast
	case len(p.opers) == 0:
		panic(ErrSyntax)
	}
	operIndex := p.opers[0]
	lTokens, rTokens := p.splitTokens(operIndex)
	p.opers = p.opers[1:]

	p.tokens = []scanner.Token{p.tokens[operIndex]}
	ast := *p.parse()

	p.tokens = lTokens
	ast.Left = p.parse()

	p.tokens = rTokens
	ast.Right = p.parse()

	return &ast
}

func (p *parser) parseVar() *Ast {
	var ast Ast
	ast.Tok = p.tokens[0]
	if p.tokens[0] == scanner.VAR {
		ast.Lit = p.vars[0]
		p.vars = p.vars[1:]
	}
	return &ast
}

func (p *parser) splitTokens(operIndex int) (lTokens []scanner.Token, rTokens []scanner.Token) {
	lTokens = p.tokens[:operIndex]
	rTokens = p.tokens[operIndex+1:]
	for i := 1; i < len(p.opers); i++ {
		p.opers[i] = p.opers[i] - operIndex - 1
	}
	return
}
