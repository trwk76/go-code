package api

import "github.com/trwk76/gocode/web/api/spec"

type (
	RequestBody interface {
		reqBodySpec() spec.RequestBodyOrRef
		reqBodyImpl() *RequestBodyImpl
	}

	RequestBodyImpl struct {
		Description string
		Required    bool
		Content     MediaTypes
	}

	RequestBodyRef struct {
		a   *API
		key string
	}

	RequestBodies struct {
		api  *API
		keys map[string]*RequestBodyImpl
	}
)

func (r *RequestBodyImpl) reqBodySpec() spec.RequestBodyOrRef {
	return spec.RequestBodyOrRef{Item: r.spec()}
}

func (r *RequestBodyImpl) reqBodyImpl() *RequestBodyImpl {
	return r
}

func (r *RequestBodyImpl) spec() spec.RequestBody {
	return spec.RequestBody{
		Description: r.Description,
		Required:    r.Required,
		Content:     r.Content.spec(),
	}
}

func (r *RequestBodyRef) reqBodySpec() spec.RequestBodyOrRef {
	return spec.RequestBodyOrRef{Ref: spec.ComponentsRef("requestBodies", r.key)}
}

func (r *RequestBodyRef) reqBodyImpl() *RequestBodyImpl {
	return r.a.RequestBodies.keys[r.key]
}

func newRequestBodies(api *API) RequestBodies {
	return RequestBodies{
		api:  api,
		keys: make(map[string]*RequestBodyImpl),
	}
}

var (
	_ RequestBody = (*RequestBodyImpl)(nil)
	_ RequestBody = (*RequestBodyRef)(nil)
)
