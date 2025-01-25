package api_test

import (
	"fmt"
	"testing"

	golang "github.com/trwk76/gocode/go"
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

	schResp = a.Schemas.Add("response", &api.StructSchema{
		Fields: []api.StructField{
			{
				Name: "corrId",
				Schema: &api.SimpleSchema{
					Spec: spec.Schema{
						Type: spec.TypeString,
						Format: spec.Format("uuid"),
					},
				},
			},
			{
				Name: "status",
				Schema: &api.SimpleSchema{
					Spec: spec.Schema{
						Type: spec.TypeInteger,
						Minimum: 200,
						Maximum: 599,
					},
				},
			},
		},
	})

	for name, body := range a.Generate(golang.PkgName("v1")) {
		fmt.Printf("%s: %s\n", name, string(body))
	}
}

var (
	schResp api.SchemaRef
	resResp api.ResponseRef
)
