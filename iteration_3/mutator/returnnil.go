package mutator

import (
	"github.com/zimmski/go-mutesting/mutator"
	"go/ast"
	"go/types"
)

func init() {
	mutator.Register("return/nil", ExampleMutatorReturnNil)
}

// MutatorReturnNil implements a mutator nilling all errors.
func ExampleMutatorReturnNil(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.ReturnStmt)
	if !ok {
		return nil
	}
	old := make([]ast.Expr, len(n.Results))
	copy(old, n.Results)
	return []mutator.Mutation{
		{
			Change: func() {
				for i, _ := range n.Results {
					n.Results[i] = ast.NewIdent("nil")
				}
			},
			Reset: func() {
				n.Results = old
			},
		},
	}
}
