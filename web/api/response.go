package api

import "github.com/trwk76/gocode/web/api/spec"

type (
	Response interface {
		respSpec() spec.ResponseOrRef
		respImpl() *ResponseImpl
	}

	ResponseImpl struct {
		Description string
		Content     MediaTypes
	}

	ResponseRef struct {
		a   *API
		key string
	}

	Responses struct {
		keys map[string]*ResponseImpl
	}
)

func (r *ResponseImpl) respSpec() spec.ResponseOrRef {
	return spec.ResponseOrRef{Item: r.spec()}
}

func (r *ResponseImpl) respImpl() *ResponseImpl {
	return r
}

func (r *ResponseImpl) spec() spec.Response {
	return spec.Response{
		Description: r.Description,
		Content:     r.Content.spec(),
	}
}

func (r *ResponseRef) respSpec() spec.ResponseOrRef {
	return spec.ResponseOrRef{Ref: spec.ComponentsRef("responses", r.key)}
}

func (r *ResponseRef) respImpl() *ResponseImpl {
	return r.a.Responses.keys[r.key]
}

func newResponses() Responses {
	return Responses{
		keys: make(map[string]*ResponseImpl),
	}
}

var (
	_ Response = (*ResponseImpl)(nil)
	_ Response = (*ResponseRef)(nil)
)
