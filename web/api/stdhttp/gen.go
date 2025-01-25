package stdhttp

import (
	"net/http"

	g "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api"
)

func NewGenerator(mapUnit *g.Unit, modelUnit *g.Unit) Generator {
	return Generator{
		mapUnit: mapUnit,
		mdlUnit: modelUnit,
	}
}

func (gen *Generator) Finalize() {
	mux := g.SymbolFor[http.ServeMux](gen.mapUnit)

	gen.mapUnit.Decls = append(
		gen.mapUnit.Decls,
		g.FuncDecls{
			g.FuncDecl{
				ID:     g.ID("Map"),
				Params: g.Params{
					{
						ID:   g.ID("m"),
						Type: g.PtrType{Item: mux},
					},
				},
				Body: gen.mapStmts,
			},
		},
	)
}

type (
	Generator struct {
		mapUnit  *g.Unit
		mapStmts g.BlockStmt
		mdlUnit  *g.Unit
		mdlTypes g.TypeDecls
	}
)

var (
	_ api.Generator = (*Generator)(nil)
)
