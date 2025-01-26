package stdhttp

import (
	"fmt"

	g "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api"
	"github.com/trwk76/gocode/web/api/spec"
)

func (gen *Generator) Operation(path any, method string, o *api.Operation, spec spec.Operation) {
	pth, ok := path.(pathHandle)
	if !ok {
		panic(fmt.Errorf("path handle expected as path; %v found", path))
	}

	val := gen.opWrap(
		g.Symbol{ID: g.ID(gen.opIDXform(spec.OperationID))},
		pth.path,
		method,
		o,
		spec,
	)

	gen.MapStmts = append(
		gen.MapStmts,
		g.ExprStmt{
			Expr: g.CallExpr{
				Func: g.MemberExpr{
					Value: varMux,
					ID:    funcHandleFunc,
				},
				Args: g.Exprs{
					g.StringExpr(fmt.Sprintf("%s %s", method, gen.opPath(gen.baseURL, pth.path))),
					val,
				},
			},
		},
	)
}

type (
	OperationPathFunc func(baseURL string, relPath string) string
	OperationWrapFunc func(expr g.Expr, path string, method string, o *api.Operation, spec spec.Operation) g.Expr
)

func defaultOperationWrapper(expr g.Expr, path string, method string, o *api.Operation, spec spec.Operation) g.Expr {
	return expr
}

func defaultOpPath(baseURL string, relPath string) string {
	return baseURL + relPath
}
