package stdhttp

import "github.com/trwk76/gocode/web/api"

func (gen *Generator) NamedPath(parent any, name string) any {
	var res pathHandle

	name = "/" + name

	if par, ok := parent.(pathHandle); ok {
		res.path = par.path + name
		res.params = par.params
	} else {
		res.path = name
	}

	return res
}

func (gen *Generator) ParamPath(parent any, name string, param api.Parameter) any {
	var res pathHandle

	name = "/" + name

	if par, ok := parent.(pathHandle); ok {
		res.path = par.path + name
		res.params = append(par.params, param)
	} else {
		res.path = name
		res.params = []api.Parameter{param}
	}

	return res
}

type (
	pathHandle struct {
		path   string
		params []api.Parameter
	}
)
