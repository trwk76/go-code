package api

import (
	"fmt"
	"slices"

	"github.com/trwk76/gocode/web/api/spec"
)

type (
	buildContext struct {
		api    *API
		gen    Generator
		ghdl   any
		path   string
		opID   string
		params []Parameter
		tags   []string
	}
)

func (c buildContext) namedChild(name string) buildContext {
	var ghdl any

	if c.gen != nil {
		ghdl = c.gen.NamedPath(c.ghdl, name)
	}

	return buildContext{
		api:    c.api,
		gen:    c.gen,
		ghdl:   ghdl,
		path:   c.path + "/" + name,
		params: c.params,
	}
}

func (c buildContext) paramChild(param Parameter) buildContext {
	var ghdl any

	pi := param.paramImpl()

	if pi.In != spec.ParameterPath {
		panic(fmt.Errorf("parameter '%s' used in path but targets '%s'", pi.Name, pi.In))
	}

	if slices.ContainsFunc(c.params, func(p Parameter) bool { return p.paramImpl().Name == pi.Name }) {
		panic(fmt.Errorf("path '%s' already defines a parameter named '%s'", c.path, pi.Name))
	}

	name := fmt.Sprintf("{%s}", pi.Name)

	if c.gen != nil {
		ghdl = c.gen.ParamPath(c.ghdl, name, pi)
	}

	return buildContext{
		api:    c.api,
		gen:    c.gen,
		ghdl:   ghdl,
		path:   c.path + "/" + name,
		params: append(c.params, param),
	}
}

func mergeTags(org []string, items []string) []string {
	for _, item := range items {
		if !slices.Contains(org, item) {
			org = append(org, item)
		}
	}

	return org
}
