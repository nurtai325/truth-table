package parser_test

import (
	"testing"

	"github.com/nurtai325/truth-table/internal/parser"
)

func TestParserNormal(t *testing.T) {
	exp := "!(a || !b) -> (b && d) <=> !b"
	testAstNodes := []*parser.Ast{}

	ast, err := parser.Parse(&exp)
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	ast.Walk(func(ast *parser.Ast) {
		if !equalAst(ast, testAstNodes[i]) {
			t.Fatalf("%d %v %v", i, ast, testAstNodes[i])
		}
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

func TestParserErr(t *testing.T) {
}
