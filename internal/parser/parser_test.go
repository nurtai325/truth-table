package parser_test

import (
	"testing"

	"github.com/nurtai325/truth-table/internal/parser"
	"github.com/nurtai325/truth-table/internal/scanner"
)

func TestParserNormal(t *testing.T) {
	exp := "!(a || !b) -> (b && d) <=> !b"
	testAstNodes := [...]*parser.Ast{
		{nil, nil, scanner.NOT, "", false, nil},
		{nil, nil, scanner.LPAREN, "", false, nil},
		{nil, nil, scanner.VAR, "a", false, nil},
		{nil, nil, scanner.OR, "", false, nil},
		{nil, nil, scanner.VAR, "b", true, nil},
		{nil, nil, scanner.RPAREN, "", false, nil},
		{nil, nil, scanner.IMPLICATION, "", false, nil},
		{nil, nil, scanner.LPAREN, "", false, nil},
		{nil, nil, scanner.VAR, "b", false, nil},
		{nil, nil, scanner.AND, "", false, nil},
		{nil, nil, scanner.VAR, "d", false, nil},
		{nil, nil, scanner.RPAREN, "", false, nil},
		{nil, nil, scanner.IF_AND_ONLY_IF, "", false, nil},
		{nil, nil, scanner.VAR, "b", true, nil},
	}
	ast, err := parser.Parse(&exp)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	ast.DfsWalk(func(ast *parser.Ast) {
		if !equalAst(testAstNodes[i], ast) {
			t.Fatalf("%d: expected: %v actual: %v", i+1, testAstNodes[i], ast)
		}
		i++
	})
}

func equalAst(a, b *parser.Ast) bool {
	switch {
	case a.Tok != b.Tok:
		return false
	case a.Lit != b.Lit:
		return false
	default:
		return true
	}
}

// TODO: test correct err checking
func TestParserErr(t *testing.T) {
}
