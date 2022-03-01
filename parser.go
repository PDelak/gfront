package main

import("go/ast"
       "go/parser"
       "go/token"
       "fmt"
)

func main() {
        fset := token.NewFileSet()
        node, err := parser.ParseFile(fset, "model.go", nil, parser.ParseComments)
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
}
