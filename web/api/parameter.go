package api

import (
	"fmt"

	"github.com/trwk76/gocode/web/api/spec"
)

func (p *Parameters) Add(key string, impl *ParameterImpl) ParameterRef {
	key = uniqueKey(p.keys, key, "param")

	res := ParameterRef{
		a:   p.api,
		key: key,
	}

	p.keys[key] = impl

	return res
}

type (
	Parameter interface {
		paramSpec() spec.ParameterOrRef
		paramImpl() *ParameterImpl
	}

	ParameterImpl struct {
		Name        string
		In          spec.ParameterIn
		Description string
		Required    bool
		Deprecated  bool
		Schema      Schema
	}

	ParameterRef struct {
		a   *API
		key string
	}

	Parameters struct {
		api  *API
		keys map[string]*ParameterImpl
	}
)

func (p *ParameterImpl) paramSpec() spec.ParameterOrRef {
	return spec.ParameterOrRef{Item: p.spec()}
}

func (p *ParameterImpl) paramImpl() *ParameterImpl {
	return p
}

func (p *ParameterImpl) spec() spec.Parameter {
	res := spec.Parameter{
		Name:        p.Name,
		In:          p.In,
		Description: p.Description,
		Required:    p.Required,
		Deprecated:  p.Deprecated,
	}

	if p.Schema != nil {
		sch := p.Schema.schemaSpec()
		s := p.Schema.schemaImpl().schema()

		switch s.Type {
		case spec.TypeBoolean, spec.TypeInteger, spec.TypeNumber, spec.TypeString, spec.TypeArray:
		default:
			panic(fmt.Errorf("only simple types can be handled by parameters"))
		}

		res.Schema = &sch
	}

	return res
}

func (r *ParameterRef) paramSpec() spec.ParameterOrRef {
	return spec.ParameterOrRef{Ref: spec.ComponentsRef("parameters", r.key)}
}

func (r *ParameterRef) paramImpl() *ParameterImpl {
	return r.a.Parameters.keys[r.key]
}

func newParameters(api *API) Parameters {
	return Parameters{
		api:  api,
		keys: make(map[string]*ParameterImpl),
	}
}

func (p Parameters) spec() spec.NamedParameterOrRefs {
	res := make(spec.NamedParameterOrRefs)

	for key, impl := range p.keys {
		res[key] = spec.ParameterOrRef{Item: impl.spec()}
	}

	return res
}

var (
	_ Parameter = (*ParameterImpl)(nil)
	_ Parameter = (*ParameterRef)(nil)
)
