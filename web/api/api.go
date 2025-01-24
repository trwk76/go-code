package api

import (
	"strings"

	"github.com/trwk76/gocode/web/api/spec"
)

func NewAPI(baseURL string) *API {
	if !strings.HasPrefix(baseURL, "/") {
		baseURL = "/" + baseURL
	}

	res := &API{
		baseURL:       baseURL,
	}

	res.Schemas = newSchemas(res)
	res.Parameters = newParameters(res)
	res.RequestBodies = newRequestBodies(res)
	res.Responses = newResponses(res)

	return res
}

type (
	API struct {
		baseURL         string
		Info            spec.Info
		Schemas         Schemas
		Parameters      Parameters
		RequestBodies   RequestBodies
		Responses       Responses
		Paths           NamedPaths
		SecuritySchemes NamedSecuritySchemes
		Security        spec.SecurityRequirements
		Tags            []spec.Tag
	}
)

func (a *API) generate() spec.OpenAPI {
	res := spec.OpenAPI{
		OpenAPI: spec.Version,
		Info:    a.Info,
		Servers: []spec.Server{{URL: a.baseURL, Description: "Current server."}},
		Paths:   make(spec.Paths),
		Components: &spec.Components{
			Schemas:         make(spec.NamedSchemas),
			Parameters:      make(spec.NamedParameterOrRefs),
			Responses:       make(spec.NamedResponseOrRefs),
			RequestBodies:   make(spec.NamedRequestBodyOrRefs),
			SecuritySchemes: make(spec.NamedSecuritySchemeOrRefs),
		},
		Security: a.Security,
		Tags:     a.Tags,
	}

	for key, item := range a.SecuritySchemes {
		res.Components.SecuritySchemes[key] = spec.SecuritySchemeOrRef{Item: item}
	}

	a.Paths.build(
		buildContext{
			api:    a,
			path:   "",
			params: nil,
			tags:   nil,
		},
		res.Paths,
	)

	return res
}
