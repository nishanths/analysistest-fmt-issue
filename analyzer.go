package analysistestissue

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// RemoveVars is an Analyzer whose suggested fix is intended to removes all top-level
// variable/constant declarations. However, it is buggy and produces broken fixes.
var RemoveVars = &analysis.Analyzer{
	Name:     "analysistestissue",
	Doc:      "shows an issue with analysistest",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder([]ast.Node{&ast.GenDecl{}}, func(n ast.Node) {
		v := n.(*ast.GenDecl)

		pass.Report(analysis.Diagnostic{
			Pos:     v.Pos(), // should not matter
			End:     v.Pos(), // should not matter
			Message: "removing",
			SuggestedFixes: []analysis.SuggestedFix{
				remove(v),
			},
		})
	})

	return nil, nil
}

func remove(v *ast.GenDecl) analysis.SuggestedFix {
	edit := analysis.TextEdit{
		Pos: v.Pos(),
		End: v.End(),
		// should produce "" to correctly remove the declaration,
		// but instead does something wrong that ends up in a broken Go file...
		NewText: []byte("<<<>>>"),
	}
	return analysis.SuggestedFix{
		Message:   "",
		TextEdits: []analysis.TextEdit{edit},
	}
}
