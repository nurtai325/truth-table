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
	Vars    map[string]bool
}

func (a *Ast) Eval() bool {
	if a.Tok == scanner.VAR {
		value := a.Vars[a.Lit]
		if a.Negated {
			value = !value
		}
		return value
	}
	a.Left.Vars = a.Vars
	a.Right.Vars = a.Vars
	l, r := a.Left.Eval(), a.Right.Eval()
	res := scanner.Operate(l, r, a.Tok)
	if a.Negated {
		res = !res
	}
	return res
}

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
	return fmt.Sprintf("%v %s negated: %t", a.Tok, a.Lit, a.Negated)
}
