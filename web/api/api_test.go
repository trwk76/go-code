package api_test

import (
	"fmt"
	"testing"

	"github.com/trwk76/gocode/web/api"
	"github.com/trwk76/gocode/testhelpers"
)

func TestAPI(t *testing.T) {
	a := api.NewAPI("api/v1")

	testhelpers.SetupAPI(a)

	spec := a.Generate(nil)
	fmt.Printf("%s\n", string(spec.YAML()))
}
