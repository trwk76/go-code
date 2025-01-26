package stdhttp

import (
	g "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api"
)

func (gen *Generator) Boolean(key string, impl *api.Boolean) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Enum(key string, impl *api.Enum) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Integer(key string, impl *api.Integer) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Uinteger(key string, impl *api.Uinteger) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Float(key string, impl *api.Float) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) String(key string, impl *api.String) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: g.String})
}

func (gen *Generator) Array(key string, impl *api.Array) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Map(key string, impl *api.Map) {
	gen.AddTypeDecl(key, g.TypeAlias{Target: convertType(gen.tcnv, impl, "")})
}

func (gen *Generator) Struct(key string, impl *api.Struct) {
	gen.AddTypeDecl(key, convertType(gen.tcnv, impl, ""))
}

func (gen *Generator) AddTypeDecl(name string, spec g.TypeSpec) {
	gen.mdlTypes = append(gen.mdlTypes, g.TypeDecl{
		ID:   g.ID(name),
		Spec: spec,
	})
}
