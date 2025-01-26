package api

import "github.com/trwk76/go-code/web/api/spec"

func (r *RequestBodies) Add(key string, impl *RequestBodyImpl) RequestBodyRef {
	key = uniqueKey(r.keys, key, "reqBody")

	res := RequestBodyRef{
		a:   r.api,
		key: key,
	}

	r.keys[key] = impl

	return res
}

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

func (r RequestBodies) spec(g Generator) spec.NamedRequestBodyOrRefs {
	res := make(spec.NamedRequestBodyOrRefs)

	for key, impl := range r.keys {
		res[key] = spec.RequestBodyOrRef{Item: impl.spec()}

		if g != nil {
			g.RequestBody(key, impl)
		}
	}

	return res
}

var (
	_ RequestBody = (*RequestBodyImpl)(nil)
	_ RequestBody = (*RequestBodyRef)(nil)
)
