package stdhttp_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	code "github.com/trwk76/go-code"
	golang "github.com/trwk76/go-code/go"
	"github.com/trwk76/go-code/testhelpers"
	"github.com/trwk76/go-code/web/api"
	"github.com/trwk76/go-code/web/api/stdhttp"
)

func TestGen(t *testing.T) {
	a := api.NewAPI("/api/test/")

	testhelpers.SetupAPI(a)

	unit := golang.Unit{
		Package: golang.PkgName("testapi"),
	}

	gen := stdhttp.NewGenerator(
		&unit,
		&unit,
		nil,
		nil,
		nil,
		nil,
	)

	uuid := golang.SymbolFor[uuid.UUID](&unit)

	gen.MdlTypes = append(gen.MdlTypes, golang.TypeDecl{
		ID: golang.ID("uuid"),
		Spec: golang.TypeAlias{
			Target: &uuid,
		},
	})

	a.Generate(&gen)

	fmt.Println(code.WriteString("\t", func(w *code.Writer) {
		unit.Write(w)
	}))
}
