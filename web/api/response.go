package api

import "github.com/trwk76/go-code/web/api/spec"

func (r *Responses) Add(key string, impl *ResponseImpl) ResponseRef {
	key = uniqueKey(r.keys, key, "response")

	res := ResponseRef{
		a:   r.api,
		key: key,
	}

	r.keys[key] = impl

	return res
}

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
		api  *API
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

func newResponses(api *API) Responses {
	return Responses{
		api:  api,
		keys: make(map[string]*ResponseImpl),
	}
}

func (r Responses) spec(g Generator) spec.NamedResponseOrRefs {
	res := make(spec.NamedResponseOrRefs)

	for key, impl := range r.keys {
		res[key] = spec.ResponseOrRef{Item: impl.spec()}

		if g != nil {
			g.Response(key, impl)
		}
	}

	return res
}

var (
	_ Response = (*ResponseImpl)(nil)
	_ Response = (*ResponseRef)(nil)
)
