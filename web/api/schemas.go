package api

import (
	golang "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api/spec"
)

type (
	Schema interface {
		schemaImpl() SchemaImpl
		schemaSpec() spec.SchemaOrRef
	}

	SchemaImpl interface {
		Schema
		schema() spec.Schema
	}

	SimpleSchema struct {
		Spec spec.Schema
		Code golang.Type
	}

	SchemaRef struct {
		a   *API
		key string
	}

	Schemas struct {
		keys map[string]SchemaImpl
	}
)

var (
	Boolean SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeBoolean},
		Code: golang.Bool,
	}

	Int32 SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeInteger, Format: spec.FormatInt32},
		Code: golang.Int32,
	}

	Int64 SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeInteger, Format: spec.FormatInt64},
		Code: golang.Int64,
	}

	Floaat SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeNumber, Format: spec.FormatFloat},
		Code: golang.Float32,
	}

	Double SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeNumber, Format: spec.FormatDouble},
		Code: golang.Float64,
	}

	String SimpleSchema = SimpleSchema{
		Spec: spec.Schema{Type: spec.TypeString},
		Code: golang.String,
	}
)

func newSchemas() Schemas {
	return Schemas{
		keys: make(map[string]SchemaImpl),
	}
}

func (r *SimpleSchema) schemaImpl() SchemaImpl {
	return r
}

func (r *SimpleSchema) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Item: r.Spec}
}

func (r *SimpleSchema) schema() spec.Schema {
	return r.Spec
}

func (r *SchemaRef) schemaImpl() SchemaImpl {
	return r.a.Schemas.keys[r.key]
}

func (r *SchemaRef) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Ref: spec.ComponentsRef("schemas", r.key)}
}

var (
	_ SchemaImpl = (*SimpleSchema)(nil)
	_ Schema     = (*SchemaRef)(nil)
)
