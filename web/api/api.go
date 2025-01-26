package api

import (
	"strings"

	"github.com/trwk76/go-code/web/api/spec"
)

func NewAPI(baseURL string) *API {
	if !strings.HasPrefix(baseURL, "/") {
		baseURL = "/" + baseURL
	}

	baseURL = strings.TrimSuffix(baseURL, "/")

	res := &API{
		baseURL: baseURL,
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

func (a *API) Generate(g Generator) spec.OpenAPI {
	if g != nil {
		g.Initialize(a.baseURL)
	}

	res := spec.OpenAPI{
		OpenAPI: spec.Version,
		Info:    a.Info,
		Servers: []spec.Server{{URL: a.baseURL, Description: "Current server."}},
		Paths:   make(spec.Paths),
		Components: &spec.Components{
			Schemas:         a.Schemas.generate(g),
			Parameters:      a.Parameters.spec(g),
			Responses:       a.Responses.spec(g),
			RequestBodies:   a.RequestBodies.spec(g),
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
			gen:    g,
			path:   "",
			params: nil,
			tags:   nil,
		},
		res.Paths,
	)

	if g != nil {
		g.Finalize(res)
	}

	return res
}
