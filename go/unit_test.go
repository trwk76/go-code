package golang_test

import (
	_ "embed"
	"encoding"
	"fmt"
	"testing"

	"github.com/google/uuid"
	code "github.com/trwk76/gocode"
	golang "github.com/trwk76/gocode/go"
)

func TestUnit(t *testing.T) {
	for _, item := range testItems {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("panic: %v", r)
				}

				t.Errorf("test '%s' failed: %s", item.name, err.Error())
			}
		}()

		unit := item.gen()
		res := code.WriteString("\t", func(w *code.Writer) { unit.Write(w) })

		if res != item.text {
			t.Errorf("test '%s' failed; expected:\n%s\ngot:\n%s\n", item.name, item.text, res)
		}
	}
}

type (
	testItem struct {
		name string
		gen  func() golang.Unit
		text string
	}
)

var testItems []testItem = []testItem{
	{
		name: "Simple",
		gen: func() golang.Unit {
			res := golang.Unit{
				Prefix:  golang.Comment(" THIS FILE IS AUTOMATICALLY GENERATED; DO NOT EDIT"),
				Package: golang.PkgName("test"),
			}

			uuid := golang.SymbolFor[uuid.UUID](&res)
			encMarsh := golang.SymbolFor[encoding.TextMarshaler](&res)
			encUnmarsh := golang.SymbolFor[encoding.TextUnmarshaler](&res)

			res.Decls = append(
				res.Decls,
				golang.TypeDecls{
					{
						ID: golang.ID("ID"),
						Spec: golang.StructType{
							Fields: []golang.StructField{
								{
									ID:   "Value",
									Type: uuid,
									Tags: golang.Tags{
										{
											Name:  "json",
											Value: "value",
										},
									},
								},
							},
						},
					},
				},
				golang.VarDecls{
					{
						ID:   golang.Ignore,
						Type: encMarsh,
						Value: golang.StructExpr{
							Type: golang.Symbol{
								ID: golang.ID("ID"),
							},
						},
					},
					{
						ID:   golang.Ignore,
						Type: encUnmarsh,
						Value: golang.CastExpr{
							Type: golang.PtrType{
								Item: golang.Symbol{
									ID: golang.ID("ID"),
								},
							},
							Value: golang.Nil,
						},
					},
				},
			)

			return res
		},
		text: simpleText,
	},
}

//go:embed tests/simple_test.go
var simpleText string
