package testhelpers

import (
	"math"
	"net/http"

	"github.com/trwk76/go-code/web/api"
	"github.com/trwk76/go-code/web/api/spec"
)

func SetupAPI(a *api.API) {
	var (
		schResp     api.SchemaRef
		schPageResp api.SchemaRef
		schErrResp  api.SchemaRef
		schCountry  api.SchemaRef
		schUUID     api.SchemaRef
		respError   api.ResponseRef
	)

	respError = a.Responses.Add("respErr", &api.ResponseImpl{
		Description: "An error occurred",
		Content: api.MediaTypes{
			api.MediaTypeJSON: api.MediaType{
				Schema: &schErrResp,
			},
		},
	})

	schResp = a.Schemas.Add("response", &api.Struct{
		Fields: []api.StructField{
			{
				Name:   "corrId",
				Schema: &schUUID,
			},
			{
				Name: "status",
				Schema: &api.Integer{
					Minimum: 200,
					Maximum: 599,
				},
			},
		},
	})

	schErrResp = a.Schemas.Add("errorResponse", &api.Struct{
		Bases: []api.Schema{&schResp},
		Fields: []api.StructField{
			{
				Name:   "message",
				Schema: &api.String{},
			},
		},
	})

	schPageResp = a.Schemas.Add("pageResponse", &api.Struct{
		Bases: []api.Schema{&schResp},
		Fields: []api.StructField{
			{
				Name:   "totalCount",
				Schema: &api.Uinteger{},
			},
			{
				Name: "pageIndex",
				Schema: &api.Uinteger{
					Maximum: math.MaxUint,
				},
			},
			{
				Name: "pageSize",
				Schema: &api.Uinteger{
					Minimum: 1,
					Maximum: math.MaxUint32,
				},
			},
		},
	})

	a.Paths = api.NamedPaths{
		"country": api.Path{
			OperationID: "country",
			GET: &api.Operation{
				ID: "Search",
				Responses: api.ResponseMap{
					Codes: map[int]api.Response{
						http.StatusOK: &api.ResponseImpl{
							Description: "Page of country results",
							Content: api.MediaTypes{
								api.MediaTypeJSON: api.MediaType{
									Schema: &api.Struct{
										Bases: []api.Schema{&schPageResp},
										Fields: []api.StructField{
											{
												Name: "items",
												Schema: &api.Array{
													Items: &schCountry,
												},
											},
										},
									},
								},
							},
						},
					},
					Default: &respError,
				},
			},
		},
	}

	schCountry = a.Schemas.Add("country", &api.Struct{
		Fields: []api.StructField{
			{
				Name: "iso3166a2",
				Schema: &api.String{
					MinLength: 2,
					MaxLength: 2,
					Pattern:   "^[A-Z]{2}$",
				},
			},
			{
				Name: "iso3166a3",
				Schema: &api.String{
					MinLength: 3,
					MaxLength: 3,
					Pattern:   "^[A-Z]{3}$",
				},
			},
			{
				Name:   "name",
				Schema: &api.String{},
			},
		},
	})

	schUUID = a.Schemas.Add("uuid", &api.String{
		Format: spec.Format("uuid"),
	})
}
