package jsonCameler

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	doc           = "jsonCameler is ..."
	ignoreComment = "nocamel"
)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "jsonCameler",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		st, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		for _, field := range st.Fields.List {
			fieldName := field.Names[0]
			tag := field.Tag.Value

			for _, str := range []string{"`", "\"", "json", ":"} {
				tag = strings.ReplaceAll(tag, str, "")
			}

			for _, str := range []string{"_", "-"} {
				if strings.Contains(tag, str) {
					if len(tag) == 1 && str == "-" {
						continue
					}

					if field.Comment != nil {
						if canIgnore(field.Comment.List) {
							continue
						}
					}

					pass.Reportf(fieldName.Pos(), fmt.Sprintf("invalid JSON tag `%s`", tag))
				}
			}
		}
	})

	return nil, nil
}

func canIgnore(comments []*ast.Comment) bool {
	if comments == nil {
		return false
	}
	for _, cm := range comments {
		if strings.Contains(cm.Text, ignoreComment) {
			return true
		}
	}
	return false
}
