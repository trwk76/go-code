package testhelpers

import (
	"math"
	"net/http"

	"github.com/trwk76/go-code/web/api"
	"github.com/trwk76/go-code/web/api/spec"
)

func SetupAPI(a *api.API) {
	var (
		schResp      api.SchemaRef
		schPageResp  api.SchemaRef
		schErrResp   api.SchemaRef
		schCountry   api.SchemaRef
		schISO3166a2 api.SchemaRef
		schUUID      api.SchemaRef
		respError    api.ResponseRef
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
				OperationID: "Search",
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
			Param: &api.ParamPath{
				Param: &api.ParameterImpl{
					Name:     "iso3166a2",
					In:       spec.ParameterPath,
					Required: true,
					Schema:   &schISO3166a2,
				},
				Path: api.Path{
					GET: &api.Operation{
						OperationID: "Fetch",
						Summary:     "Fetch country",
						Description: "Fetch a country given its code.",
						Responses: api.ResponseMap{
							Codes: map[int]api.Response{
								http.StatusOK: &api.ResponseImpl{
									Description: "Country",
									Content: api.MediaTypes{
										api.MediaTypeJSON: api.MediaType{
											Schema: &schCountry,
										},
									},
								},
							},
							Default: &respError,
						},
					},
					Named: api.NamedPaths{
						"regions": api.Path{
							OperationID: "Regions",
							GET: &api.Operation{
								OperationID: "Search",
								Summary:     "Search country's regions",
								Description: "Search country's regions",
								Responses: api.ResponseMap{
									Codes: map[int]api.Response{
										http.StatusOK: &api.ResponseImpl{
											Description: "Country",
											Content: api.MediaTypes{
												api.MediaTypeJSON: api.MediaType{
													Schema: &schCountry,
												},
											},
										},
									},
									Default: &respError,
								},
							},
						},
					},
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
				Name:   "iso3166a3",
				Schema: &schISO3166a2,
			},
			{
				Name:   "name",
				Schema: &api.String{},
			},
		},
	})

	schISO3166a2 = a.Schemas.Add("iso3166a2", &api.String{
		MinLength: 3,
		MaxLength: 3,
		Pattern:   "^[A-Z]{3}$",
	})

	schUUID = a.Schemas.Add("uuid", &api.String{
		Format: spec.Format("uuid"),
	})
}
