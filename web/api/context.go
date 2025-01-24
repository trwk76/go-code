package api

import (
	"fmt"
	"slices"

	"github.com/trwk76/gocode/web/api/spec"
)

type (
	buildContext struct {
		api    *API
		path   string
		params []Parameter
		tags   []string
	}
)

func (c buildContext) namedChild(name string) buildContext {
	return buildContext{
		api:    c.api,
		path:   c.path + "/" + name,
		params: c.params,
	}
}

func (c buildContext) paramChild(param Parameter) buildContext {
	pi := param.paramImpl()

	if pi.In != spec.ParameterPath {
		panic(fmt.Errorf("parameter '%s' used in path but targets '%s'", pi.Name, pi.In))
	}

	if slices.ContainsFunc(c.params, func(p Parameter) bool { return p.paramImpl().Name == pi.Name }) {
		panic(fmt.Errorf("path '%s' already defines a parameter named '%s'", c.path, pi.Name))
	}

	name := fmt.Sprintf("{%s}", pi.Name)

	return buildContext{
		api:    c.api,
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
