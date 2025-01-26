package api_test

import (
	"fmt"
	"testing"

	"github.com/trwk76/go-code/testhelpers"
	"github.com/trwk76/go-code/web/api"
)

func TestAPI(t *testing.T) {
	a := api.NewAPI("api/v1")

	testhelpers.SetupAPI(a)

	spec := a.Generate(nil)
	fmt.Printf("%s\n", string(spec.YAML()))
}
