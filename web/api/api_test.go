package api_test

import (
	"fmt"
	"testing"

	"github.com/trwk76/gocode/web/api"
	"github.com/trwk76/gocode/web/api/spec"
)

func TestAPI(t *testing.T) {
	a := api.NewAPI("api/v1")

	resResp = a.Responses.Add("response", &api.ResponseImpl{
		Description: "Base response",
		Content: api.MediaTypes{
			api.MediaTypeJSON: api.MediaType{
				Schema: &schResp,
			},
		},
	})

	schResp = a.Schemas.Add("response", &api.Struct{
		Fields: []api.StructField{
			{
				Name: "corrId",
				Schema: &api.String{
					Format: spec.Format("uuid"),
				},
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

	spec := a.Generate(nil)
	fmt.Printf("%s\n", string(spec.YAML()))
}

var (
	schResp api.SchemaRef
	resResp api.ResponseRef
)
