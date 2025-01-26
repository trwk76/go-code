package stdhttp

import (
	"net/http"
	"reflect"
	"sort"

	code "github.com/trwk76/go-code"
	g "github.com/trwk76/go-code/go"
	"github.com/trwk76/go-code/web/api"
	"github.com/trwk76/go-code/web/api/spec"
)

func NewGenerator(mapUnit *g.Unit, modelUnit *g.Unit, opIDXform code.IDTransformer, opPath OperationPathFunc, opWrapper OperationWrapFunc, typeConv TypeConverter) Generator {
	if opIDXform == nil {
		opIDXform = func(id string) string { return id }
	}

	if opPath == nil {
		opPath = defaultOpPath
	}

	if opWrapper == nil {
		opWrapper = defaultOperationWrapper
	}

	if typeConv == nil {
		typeConv = (*DefaultTypeConverter)(nil)
	}

	return Generator{
		mapUnit:   mapUnit,
		mdlUnit:   modelUnit,
		opIDXform: opIDXform,
		opPath:    opPath,
		opWrap:    opWrapper,
		tcnv:      reflect.TypeOf(typeConv).Elem(),
	}
}

func (gen *Generator) Initialize(baseURL string) {
	gen.baseURL = baseURL
}

func (gen *Generator) Finalize(spec spec.OpenAPI) {
	sort.Slice(gen.MdlTypes, func(i, j int) bool {
		return gen.MdlTypes[i].ID < gen.MdlTypes[j].ID
	})

	if gen.mapUnit != nil {
		mux := g.SymbolFor[http.ServeMux](gen.mapUnit)

		gen.mapUnit.Decls = append(
			gen.mapUnit.Decls,
			g.FuncDecls{
				g.FuncDecl{
					ID: g.ID("Map"),
					Params: g.Params{{
						ID:   varMux.ID,
						Type: g.PtrType{Item: mux},
					}},
					Body: gen.MapStmts,
				},
			},
		)
	}

	if gen.mdlUnit != nil {
		gen.mdlUnit.Decls = append(
			gen.mdlUnit.Decls,
			gen.MdlTypes,
		)
	}
}

type (
	Generator struct {
		baseURL   string
		mapUnit   *g.Unit
		mdlUnit   *g.Unit
		opIDXform code.IDTransformer
		opPath    OperationPathFunc
		opWrap    OperationWrapFunc
		tcnv      reflect.Type

		MapStmts g.BlockStmt
		MdlTypes g.TypeDecls
	}
)

var (
	varMux         g.Symbol = g.Symbol{ID: g.ID("mux")}
	funcHandleFunc g.ID     = g.ID("HandleFunc")
)

var (
	_ api.Generator = (*Generator)(nil)
)
