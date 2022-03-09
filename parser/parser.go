package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(file string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println(fn.Name.Name)
	}
	ast.Inspect(node, func(n ast.Node) bool {
		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			var exported string
			if fn.Name.IsExported() {
				exported = "exported "
			}
			fmt.Printf("%sfunction declaration found on line %d: \n\t%s\n", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			return true
		}
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			fmt.Println("FuncCall: ", funcCall.Fun)
		}
		return true
	})
	fmt.Println()
}
