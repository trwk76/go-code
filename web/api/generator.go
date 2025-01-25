package api

import "github.com/trwk76/gocode/web/api/spec"

type (
	Generator interface {
		Schema(name string, spec spec.Schema)
		Parameter(name string, spec spec.Parameter)
		RequestBody(name string, spec spec.RequestBody)
		Response(name string, spec spec.Response)

		NamedPath(parent any, name string) any
		ParamPath(parent any, name string, param spec.Parameter) any
		Operation(path any, method string, op spec.Operation)
	}

	MultiGenerator []Generator
	multiPath      []any
)

func (m MultiGenerator) Schema(name string, spec spec.Schema) {
	m.each(func(idx int, g Generator) { g.Schema(name, spec) })
}

func (m MultiGenerator) Parameter(name string, spec spec.Parameter) {
	m.each(func(idx int, g Generator) { g.Parameter(name, spec) })
}

func (m MultiGenerator) RequestBody(name string, spec spec.RequestBody) {
	m.each(func(idx int, g Generator) { g.RequestBody(name, spec) })
}

func (m MultiGenerator) Response(name string, spec spec.Response) {
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

func (m MultiGenerator) ParamPath(parent any, name string, param spec.Parameter) any {
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

func (m MultiGenerator) Operation(path any, method string, op spec.Operation) {
	pth := path.(multiPath)

	m.each(func(idx int, g Generator) {
		var p any

		if idx < len(pth) {
			p = pth[idx]
		}

		g.Operation(p, method, op)
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
