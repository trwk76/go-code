package stdhttp

import (
	"slices"

	g "github.com/trwk76/go-code/go"
	"github.com/trwk76/go-code/web/api"
	"github.com/trwk76/go-code/web/api/spec"
)

func (gen *Generator) Boolean(key string, impl *api.Boolean) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Enum(key string, impl *api.Enum) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Integer(key string, impl *api.Integer) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Uinteger(key string, impl *api.Uinteger) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Float(key string, impl *api.Float) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) String(key string, impl *api.String) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: g.String}, impl.Spec())
}

func (gen *Generator) Array(key string, impl *api.Array) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Map(key string, impl *api.Map) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")}, impl.Spec())
}

func (gen *Generator) Struct(key string, impl *api.Struct) {
	gen.AddTypeDecl(key, convertType(gen.tcnv, impl, ""), impl.Spec())
}

func (gen *Generator) AddTypeDecl(name string, tspec g.TypeSpec, sspec spec.Schema) {
	if slices.ContainsFunc(gen.MdlTypes, func(t g.TypeDecl) bool { return t.ID == g.ID(name) }) {
		return
	}

	gen.MdlTypes = append(gen.MdlTypes, g.TypeDecl{
		ID:   g.ID(name),
		Spec: tspec,
	})
}
