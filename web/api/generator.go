package api

import "github.com/trwk76/gocode/web/api/spec"

type (
	Generator interface {
		Initialize(baseURL string)

		Boolean(name string, spec *Boolean)
		Enum(name string, spec *Enum)
		Integer(name string, spec *Integer)
		Uinteger(name string, spec *Uinteger)
		Float(name string, spec *Float)
		String(name string, spec *String)
		Array(name string, spec *Array)
		Map(name string, spec *Map)
		Struct(name string, spec *Struct)
		Parameter(name string, spec *ParameterImpl)
		RequestBody(name string, spec *RequestBodyImpl)
		Response(name string, spec *ResponseImpl)

		NamedPath(parent any, name string) any
		ParamPath(parent any, name string, param Parameter) any
		Operation(path any, method string, op *Operation, spec spec.Operation)
	}

	MultiGenerator []Generator
	multiPath      []any
)

func (m MultiGenerator) Initialize(baseURL string) {
	m.each(func(idx int, g Generator) { g.Initialize(baseURL) })
}

func (m MultiGenerator) Boolean(name string, spec *Boolean) {
	m.each(func(idx int, g Generator) { g.Boolean(name, spec) })
}

func (m MultiGenerator) Enum(name string, spec *Enum) {
	m.each(func(idx int, g Generator) { g.Enum(name, spec) })
}

func (m MultiGenerator) Integer(name string, spec *Integer) {
	m.each(func(idx int, g Generator) { g.Integer(name, spec) })
}

func (m MultiGenerator) Uinteger(name string, spec *Uinteger) {
	m.each(func(idx int, g Generator) { g.Uinteger(name, spec) })
}

func (m MultiGenerator) Float(name string, spec *Float) {
	m.each(func(idx int, g Generator) { g.Float(name, spec) })
}

func (m MultiGenerator) String(name string, spec *String) {
	m.each(func(idx int, g Generator) { g.String(name, spec) })
}

func (m MultiGenerator) Array(name string, spec *Array) {
	m.each(func(idx int, g Generator) { g.Array(name, spec) })
}

func (m MultiGenerator) Map(name string, spec *Map) {
	m.each(func(idx int, g Generator) { g.Map(name, spec) })
}

func (m MultiGenerator) Struct(name string, spec *Struct) {
	m.each(func(idx int, g Generator) { g.Struct(name, spec) })
}

func (m MultiGenerator) Parameter(name string, spec *ParameterImpl) {
	m.each(func(idx int, g Generator) { g.Parameter(name, spec) })
}

func (m MultiGenerator) RequestBody(name string, spec *RequestBodyImpl) {
	m.each(func(idx int, g Generator) { g.RequestBody(name, spec) })
}

func (m MultiGenerator) Response(name string, spec *ResponseImpl) {
	m.each(func(idx int, g Generator) { g.Response(name, spec) })
}

func (m MultiGenerator) NamedPath(parent any, name string) any {
	res := make(multiPath, len(m))
	par := parent.(multiPath)

	m.each(func(idx int, g Generator) {
		var p any

		if idx < len(par) {
			p = par[idx]
		}

		res[idx] = g.NamedPath(p, name)
	})

	return res
}

func (m MultiGenerator) ParamPath(parent any, name string, param Parameter) any {
	res := make(multiPath, len(m))
	par := parent.(multiPath)

	m.each(func(idx int, g Generator) {
		var p any

		if idx < len(par) {
			p = par[idx]
		}

		res[idx] = g.ParamPath(p, name, param)
	})

	return res
}

func (m MultiGenerator) Operation(path any, method string, op *Operation, spec spec.Operation) {
	pth := path.(multiPath)

	m.each(func(idx int, g Generator) {
		var p any

		if idx < len(pth) {
			p = pth[idx]
		}

		g.Operation(p, method, op, spec)
	})
}

func (m MultiGenerator) each(f func(idx int, g Generator)) {
	for idx, g := range m {
		f(idx, g)
	}
}

var (
	_ Generator = MultiGenerator{}
)
