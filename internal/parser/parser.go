package parser

import (
	"errors"
	"slices"

	"github.com/nurtai325/truth-table/internal/scanner"
)

const (
	T = "T"
	F = "F"
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

func Parse(exp *string, debug bool) (t truthTable, err error) {
	defer func() {
		if debug {
			return
		}
		if e := recover(); e != nil {
			t = nil
			err = e.(error)
		}
	}()

	var ast Ast
	p := new(parser)

	tokens, operators, variables := scanAll(exp)
	p.init(tokens, operators, variables)
	vars := make([]string, len(p.vars))
	copy(vars, p.vars)
	ast = *p.parse()

	vars = p.filterVars(vars)
	varMap := make(map[string]bool, len(vars))
	values, height := emptyTable(len(vars))
	height += 1
	width := len(vars) + 1

	t = make(truthTable, height)
	header := make([]string, width)
	for i := 0; i < len(vars); i++ {
		header[i] = vars[i]
	}
	header[len(header)-1] = *exp
	t[0] = header

	for i := 0; i < height-1; i++ {
		for j := 0; j < len(vars); j++ {
			varMap[vars[j]] = values[i][j]
		}
		ast.Vars = varMap
		res := ast.Eval()

		row := make([]string, width)
		cell := F
		for j := 0; j < width-1; j++ {
			if values[i][j] {
				cell = T
			} else {
				cell = F
			}
			row[j] = cell
		}
		if res {
			cell = T
		} else {
			cell = F
		}
		row[width-1] = cell
		t[i+1] = row
	}

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
	p.insertParentheses()
}

func (p *parser) insertParentheses() {
	newTokens := make([]scanner.Token, 0, len(p.tokens))
	for i := 0; i < len(p.tokens); i++ {
		newTokens = append(newTokens, p.tokens[i])
	}
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

	negated := false
	if p.tokens[0] == scanner.NOT && p.tokens[1] == scanner.LPAREN {
		for i := 0; i < len(p.opers); i++ {
			p.opers[i] = p.opers[i] - 1
		}
		p.tokens = p.tokens[1:]
		negated = true
	}

	operIndex := p.opers[0]
	lTokens, rTokens, oper := p.splitTokens(operIndex)

	var ast Ast
	p.tokens = []scanner.Token{oper}
	if oper == scanner.EOF {
		p.tokens = lTokens
		ast = *p.parse()
		if negated {
			ast.Negated = negated
		}
		return &ast
	}

	ast = *p.parse()
	p.tokens = lTokens
	ast.Left = p.parse()
	if negated {
		ast.Left.Negated = negated
	}

	p.tokens = rTokens
	ast.Right = p.parse()

	return &ast
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

func (p *parser) splitTokens(operIndex int) ([]scanner.Token, []scanner.Token, scanner.Token) {
	var lTokens []scanner.Token
	var rTokens []scanner.Token

	if operIndex >= len(p.tokens) {
		panic(ErrSyntax)
	}
	oper := p.tokens[operIndex]
	if oper == scanner.RPAREN || oper == scanner.NOT {
		panic(ErrSyntax)
	}

	if oper != scanner.LPAREN {
		lTokens = p.tokens[:operIndex]
		rTokens = p.tokens[operIndex+1:]

		for i := 1; i < len(p.opers); i++ {
			p.opers[i] = p.opers[i] - operIndex - 1
		}
		p.opers = p.opers[1:]

		return lTokens, rTokens, oper
	}

	rParenIdx := p.nextRParen()
	rParen := p.opers[rParenIdx]
	operLen := len(p.opers)

	lTokens = p.tokens[1:rParen]
	if rParen+1 == len(p.tokens) {
		for i := 0; i < len(p.opers); i++ {
			p.opers[i] = p.opers[i] - 1
		}
		p.opers = p.opers[1 : operLen-1]
		return lTokens, rTokens, scanner.EOF
	}

	oper = p.tokens[rParen+1]
	rTokens = p.tokens[rParen+2:]

	lOpers := p.opers[:rParenIdx]
	rOpers := p.opers[rParenIdx+2:]
	for i := 0; i < len(lOpers); i++ {
		lOpers[i] = lOpers[i] - 1
	}

	offset := len(lOpers) + 2
	for i := 0; i < len(rOpers); i++ {
		rOpers[i] = rOpers[i] - offset
	}

	p.opers = slices.Concat(lOpers, rOpers)
	p.opers = p.opers[1:]
	return lTokens, rTokens, oper
}

func (p *parser) nextRParen() int {
	depth := 0
	for i := 1; i < len(p.opers); i++ {
		tokIdx := p.opers[i]
		if p.tokens[tokIdx] == scanner.LPAREN {
			depth++
			continue
		}
		if p.tokens[tokIdx] == scanner.RPAREN {
			if depth == 0 {
				return i
			}
			depth--
		}
	}
	panic(ErrSyntax)
}

func (p *parser) filterVars(unFilteredVars []string) []string {
	vars := make([]string, 0)
Outer:
	for i := 0; i < len(unFilteredVars); i++ {
		for j := 0; j < len(vars); j++ {
			if unFilteredVars[i] == vars[j] {
				continue Outer
			}
		}
		vars = append(vars, unFilteredVars[i])
	}
	return vars
}
