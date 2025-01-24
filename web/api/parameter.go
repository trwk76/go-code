package api

import "github.com/trwk76/gocode/web/api/spec"

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

func newParameters() Parameters {
	return Parameters{
		keys: make(map[string]*ParameterImpl),
	}
}

var (
	_ Parameter = (*ParameterImpl)(nil)
	_ Parameter = (*ParameterRef)(nil)
)
