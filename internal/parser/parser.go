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
	defer func() {
		if e := recover(); e != nil {
			ast = nil
			err = e.(error)
		}
	}()
	tokens, operators, variables := scanAll(exp)
	p := new(parser)
	p.init(tokens, operators, variables)
	ast = p.parse()
	ast.VarMap = p.varMap()
	return
}

func scanAll(exp *string) ([]scanner.Token, []int, []string) {
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
			panic(err)
		}

		switch token {
		case scanner.EOF:
			break Loop
		case scanner.ILLEGAL:
			panic(ErrIllegalToken)
		case scanner.NOT:
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
	return tokens, operators, variables
}

func (p *parser) init(tokens []scanner.Token, opers []int, vars []string) {
	p.tokens = tokens
	p.opers = opers
	p.vars = vars
	p.cursor = 0
}

func (p *parser) parse() *Ast {
	switch {
	case len(p.tokens) == 0:
		return nil
	case len(p.tokens) == 1:
		return p.parseToken()
	case len(p.tokens) == 2 && p.tokens[0] == scanner.NOT:
		p.tokens = p.tokens[1:]
		ast := p.parseToken()
		ast.Negated = true
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

	p.check(&ast)
	return &ast
}

func (p *parser) check(ast *Ast) {
	nodes := []*Ast{}
	parentheses := 0
	ast.DfsWalk(func(ast *Ast) {
		if len(nodes) == 0 {
			return
		}
		prevNode := nodes[len(nodes)-1]
		err := fmt.Errorf("%w at %v %v", ErrSyntax, prevNode, ast)
		if scanner.IsOperator(prevNode.Tok) && scanner.IsOperator(ast.Tok) {
			panic(err)
		} else if prevNode.Tok == scanner.VAR && ast.Tok == scanner.VAR {
			panic(err)
		}
		if ast.Tok == scanner.LPAREN {
			parentheses++
		} else if ast.Tok == scanner.RPAREN {
			parentheses--
		}
		nodes = append(nodes, ast)
	})
	if parentheses > 0 {
		panic(fmt.Errorf("%w: unclosed parentheses", ErrSyntax))
	} else if parentheses < 0 {
		panic(fmt.Errorf("%w: not opened parentheses", ErrSyntax))
	}
}

func (p *parser) parseToken() *Ast {
	var ast Ast
	ast.Tok = p.tokens[0]
	if p.tokens[0] == scanner.VAR {
		ast.Lit = p.vars[0]
		p.vars = p.vars[1:]
	}
	return &ast
}

func (p *parser) splitTokens(operIndex int) (lTokens []scanner.Token, rTokens []scanner.Token) {
	if operIndex >= len(p.tokens) {
		return make([]scanner.Token, 0), make([]scanner.Token, 0)
	}
	lTokens = p.tokens[:operIndex]
	rTokens = p.tokens[operIndex+1:]
	for i := 1; i < len(p.opers); i++ {
		p.opers[i] = p.opers[i] - operIndex - 1
	}
	return
}

func (p *parser) varMap() map[string]string {
	visited := make([]string, 0)
	variablesMap := make(map[string]string)
Outer:
	for i := 0; i < len(p.vars); i++ {
		for j := 0; i < len(visited); j++ {
			if p.vars[i] == visited[j] {
				continue Outer
			}
		}
		visited = append(visited, p.vars[i])
		variablesMap[p.vars[i]] = p.vars[i]
	}
	return variablesMap
}
