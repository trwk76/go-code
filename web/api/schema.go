package api

import (
	"errors"

	golang "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api/spec"
)

func (s *Schemas) Add(key string, impl SchemaImpl) SchemaRef {
	key = uniqueKey(s.keys, key, "schema")

	res := SchemaRef{
		a:   s.api,
		key: key,
	}

	s.keys[key] = impl

	return res
}

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

	ArraySchema struct {
		Items    Schema
		MinItems uint64
		MaxItems uint64
		Unique   bool
		Code     golang.Type
	}

	MapSchema struct {
		Key   Schema
		Value Schema
	}

	StructSchema struct {
		Bases  []Schema
		Fields []StructField
	}

	StructField struct {
		Name     string
		Schema   Schema
		Optional bool
		Code     golang.ID
	}

	SchemaRef struct {
		a   *API
		key string
	}

	Schemas struct {
		api  *API
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

func newSchemas(api *API) Schemas {
	return Schemas{
		api:  api,
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

func (s *ArraySchema) schemaImpl() SchemaImpl {
	return s
}

func (s *ArraySchema) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Item: s.schema() }
}

func (s *ArraySchema) schema() spec.Schema {
	items := s.Items.schemaSpec()

	return spec.Schema{
		Type:        spec.TypeArray,
		Items:       &items,
		MinItems:    s.MinItems,
		MaxItems:    s.MaxItems,
		UniqueItems: s.Unique,
	}
}

func (s *MapSchema) schemaImpl() SchemaImpl {
	return s
}

func (s *MapSchema) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Item: s.schema() }
}

func (s *MapSchema) schema() spec.Schema {
	res := spec.Schema{
		Type: spec.TypeObject,
	}

	key := s.Key.schemaImpl().schema()
	val := s.Value.schemaSpec()

	if key.Type != spec.TypeString {
		panic(errors.New("map key type must be a string"))
	}

	if key.Pattern != "" {
		res.Properties = spec.NamedSchemaOrRefs{
			key.Pattern: val,
		}
	} else {
		res.AdditionalProperties = &val
	}

	return res
}

func (s *StructSchema) schemaImpl() SchemaImpl {
	return s
}

func (s *StructSchema) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{ Item: s.schema() }
}

func (s *StructSchema) schema() spec.Schema {
	res := spec.Schema{
		Type:       spec.TypeObject,
		Properties: make(spec.NamedSchemaOrRefs),
	}

	for _, fld := range s.Fields {
		res.Properties[fld.Name] = fld.Schema.schemaSpec()

		if !fld.Optional {
			res.Required = append(res.Required, fld.Name)
		}
	}

	if len(s.Bases) > 0 {
		items := make([]spec.SchemaOrRef, len(s.Bases) + 1)

		for idx, base := range s.Bases {
			items[idx] = base.schemaSpec()
		}

		items[len(s.Bases)] = spec.SchemaOrRef{Item: res}
		res = spec.Schema{AllOf: items}
	}

	return res
}

func (r *SchemaRef) schemaImpl() SchemaImpl {
	return r.a.Schemas.keys[r.key]
}

func (r *SchemaRef) schemaSpec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Ref: spec.ComponentsRef("schemas", r.key)}
}

var (
	_ SchemaImpl = (*SimpleSchema)(nil)
	_ SchemaImpl = (*ArraySchema)(nil)
	_ SchemaImpl = (*MapSchema)(nil)
	_ SchemaImpl = (*StructSchema)(nil)
	_ Schema     = (*SchemaRef)(nil)
)
