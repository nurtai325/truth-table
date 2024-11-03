package parser

import (
	"container/list"
	"fmt"

	"github.com/nurtai325/truth-table/internal/scanner"
)

type Ast struct {
	Left    *Ast
	Right   *Ast
	Tok     scanner.Token
	Lit     string
	Negated bool
	VarMap  map[string]string
}

func (ast *Ast) Eval(subExpressions bool) (truthTable, error) {
	return nil, nil
}

// TODO: ast evaluation to result table
func (ast *Ast) DfsWalk(f func(ast *Ast)) {
	if ast == nil {
		return
	}
	ast.Left.DfsWalk(f)
	f(ast)
	ast.Right.DfsWalk(f)
}

func (ast *Ast) Print() {
	if ast == nil {
		return
	}

	queue := list.New()
	queue.PushBack(&nodeWithDepth{Node: ast, Depth: 0})

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		nodeDepth := element.Value.(*nodeWithDepth)
		node := nodeDepth.Node
		depth := nodeDepth.Depth

		fmt.Printf("Depth: %d, Lit: %s, Tok: %v, Negated: %v\n", depth, node.Lit, node.Tok, node.Negated)

		if node.Left != nil {
			queue.PushBack(&nodeWithDepth{Node: node.Left, Depth: depth + 1})
		}
		if node.Right != nil {
			queue.PushBack(&nodeWithDepth{Node: node.Right, Depth: depth + 1})
		}
	}
}

type nodeWithDepth struct {
	Node  *Ast
	Depth int
}

func (a Ast) String() string {
	if scanner.IsOperator(a.Tok) {
		return fmt.Sprintf("%v", a.Tok)
	}
	return fmt.Sprintf("%v %s negated: %t", a.Tok, a.Lit, a.Negated)
}
