package api

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/trwk76/go-code/web/api/spec"
)

func NewEnum[T comparable](values ...T) Enum {
	if len(values) < 1 {
		panic(fmt.Errorf("enumeration must define at least one value"))
	}

	vals := make([]any, len(values))

	for idx, val := range values {
		vals[idx] = val
	}

	return Enum{
		Type:   reflect.TypeFor[T](),
		Values: vals,
	}
}

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
		Impl() SchemaImpl
		spec() spec.SchemaOrRef
	}

	SchemaImpl interface {
		Schema
		Accept(v SchemaImplVisitor)
		Spec() spec.Schema
	}

	SchemaImplVisitor interface {
		VisitBoolean(i *Boolean)
		VisitEnum(i *Enum)
		VisitInteger(i *Integer)
		VisitUinteger(i *Uinteger)
		VisitFloat(i *Float)
		VisitString(i *String)
		VisitArray(i *Array)
		VisitMap(i *Map)
		VisitStruct(i *Struct)
	}

	Boolean struct{}

	Enum struct {
		Type   reflect.Type
		Values []any
	}

	Integer struct {
		Format     spec.Format
		Minimum    int64
		Maximum    int64
		MultipleOf int64
	}

	Uinteger struct {
		Format     spec.Format
		Minimum    uint64
		Maximum    uint64
		MultipleOf uint64
	}

	Float struct {
		Format           spec.Format
		Minimum          float64
		MinimumExclusive bool
		Maximum          float64
		MaximumExclusive bool
		MultipleOf       float64
	}

	String struct {
		Format    spec.Format
		MinLength uint64
		MaxLength uint64
		Pattern   string
	}

	Array struct {
		Items    Schema
		MinItems uint64
		MaxItems uint64
		Unique   bool
	}

	Map struct {
		Key   Schema
		Value Schema
	}

	Struct struct {
		Bases  []Schema
		Fields []StructField
	}

	StructField struct {
		Name     string
		Schema   Schema
		Optional bool
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

func (i *Boolean) Impl() SchemaImpl  { return i }
func (i *Enum) Impl() SchemaImpl     { return i }
func (i *Integer) Impl() SchemaImpl  { return i }
func (i *Uinteger) Impl() SchemaImpl { return i }
func (i *Float) Impl() SchemaImpl    { return i }
func (i *String) Impl() SchemaImpl   { return i }
func (i *Array) Impl() SchemaImpl    { return i }
func (i *Map) Impl() SchemaImpl      { return i }
func (i *Struct) Impl() SchemaImpl   { return i }

func (i *Boolean) Accept(v SchemaImplVisitor)  { v.VisitBoolean(i) }
func (i *Enum) Accept(v SchemaImplVisitor)     { v.VisitEnum(i) }
func (i *Integer) Accept(v SchemaImplVisitor)  { v.VisitInteger(i) }
func (i *Uinteger) Accept(v SchemaImplVisitor) { v.VisitUinteger(i) }
func (i *Float) Accept(v SchemaImplVisitor)    { v.VisitFloat(i) }
func (i *String) Accept(v SchemaImplVisitor)   { v.VisitString(i) }
func (i *Array) Accept(v SchemaImplVisitor)    { v.VisitArray(i) }
func (i *Map) Accept(v SchemaImplVisitor)      { v.VisitMap(i) }
func (i *Struct) Accept(v SchemaImplVisitor)   { v.VisitStruct(i) }

func (r *Boolean) Spec() spec.Schema {
	return spec.Schema{Type: spec.TypeBoolean}
}

func (i *Enum) Spec() spec.Schema {
	return spec.Schema{Enum: i.Values}
}

func (i *Integer) Spec() spec.Schema {
	res := spec.Schema{
		Type:    spec.TypeInteger,
		Format:  i.Format,
		Minimum: i.Minimum,
		Maximum: i.Maximum,
	}

	if i.MultipleOf != 0 {
		res.MultipleOf = i.MultipleOf
	}

	return res
}

func (i *Uinteger) Spec() spec.Schema {
	res := spec.Schema{
		Type:    spec.TypeInteger,
		Format:  i.Format,
		Minimum: i.Minimum,
		Maximum: i.Maximum,
	}

	if i.MultipleOf != 0 {
		res.MultipleOf = i.MultipleOf
	}

	return res
}

func (i *Float) Spec() spec.Schema {
	res := spec.Schema{
		Type:   spec.TypeNumber,
		Format: i.Format,
	}

	if i.MinimumExclusive {
		res.ExclusiveMinimum = i.Minimum
	} else {
		res.Minimum = i.Minimum
	}

	if i.MaximumExclusive {
		res.ExclusiveMaximum = i.Maximum
	} else {
		res.Maximum = i.Maximum
	}

	if i.MultipleOf != 0 {
		res.MultipleOf = i.MultipleOf
	}

	return res
}

func (i *String) Spec() spec.Schema {
	return spec.Schema{
		Type:      spec.TypeString,
		Format:    i.Format,
		MinLength: i.MinLength,
		MaxLength: i.MaxLength,
		Pattern:   i.Pattern,
	}
}

func (i *Array) Spec() spec.Schema {
	items := i.Items.spec()

	return spec.Schema{
		Type:        spec.TypeArray,
		Items:       &items,
		MinItems:    i.MinItems,
		MaxItems:    i.MaxItems,
		UniqueItems: i.Unique,
	}
}

func (i *Map) Spec() spec.Schema {
	res := spec.Schema{
		Type: spec.TypeObject,
	}

	key := i.Key.Impl().Spec()
	val := i.Value.spec()

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

func (i *Struct) Spec() spec.Schema {
	res := spec.Schema{
		Type:       spec.TypeObject,
		Properties: make(spec.NamedSchemaOrRefs),
	}

	for _, fld := range i.Fields {
		res.Properties[fld.Name] = fld.Schema.spec()

		if !fld.Optional {
			res.Required = append(res.Required, fld.Name)
		}
	}

	if len(i.Bases) > 0 {
		items := make([]spec.SchemaOrRef, len(i.Bases)+1)

		for idx, base := range i.Bases {
			items[idx] = base.spec()
		}

		items[len(i.Bases)] = spec.SchemaOrRef{Item: res}
		res = spec.Schema{AllOf: items}
	}

	return res
}

func (i *Boolean) spec() spec.SchemaOrRef  { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *Enum) spec() spec.SchemaOrRef     { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *Integer) spec() spec.SchemaOrRef  { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *Uinteger) spec() spec.SchemaOrRef { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *Float) spec() spec.SchemaOrRef    { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *String) spec() spec.SchemaOrRef   { return spec.SchemaOrRef{Item: i.Spec()} }
func (i *Array) spec() spec.SchemaOrRef    { return spec.SchemaOrRef{Item: i.Spec()} }
func (s *Map) spec() spec.SchemaOrRef      { return spec.SchemaOrRef{Item: s.Spec()} }
func (s *Struct) spec() spec.SchemaOrRef   { return spec.SchemaOrRef{Item: s.Spec()} }

func (r *SchemaRef) Key() string {
	return r.key
}

func (r *SchemaRef) Impl() SchemaImpl {
	return r.a.Schemas.keys[r.key]
}

func (r *SchemaRef) spec() spec.SchemaOrRef {
	return spec.SchemaOrRef{Ref: spec.ComponentsRef("schemas", r.key)}
}

func newSchemas(api *API) Schemas {
	return Schemas{
		api:  api,
		keys: make(map[string]SchemaImpl),
	}
}

func (s Schemas) generate(g Generator) spec.NamedSchemas {
	res := make(spec.NamedSchemas)

	for key, impl := range s.keys {
		s := impl.Spec()

		res[key] = s

		if g != nil {
			switch ti := impl.(type) {
			case *Boolean:
				g.Boolean(key, ti)
			case *Enum:
				g.Enum(key, ti)
			case *Integer:
				g.Integer(key, ti)
			case *Float:
				g.Float(key, ti)
			case *String:
				g.String(key, ti)
			case *Array:
				g.Array(key, ti)
			case *Map:
				g.Map(key, ti)
			case *Struct:
				g.Struct(key, ti)
			default:
				panic(fmt.Errorf("unsupported schema impl %T", impl))
			}
		}
	}

	return res
}

var (
	_ SchemaImpl = (*Boolean)(nil)
	_ SchemaImpl = (*Enum)(nil)
	_ SchemaImpl = (*Integer)(nil)
	_ SchemaImpl = (*Uinteger)(nil)
	_ SchemaImpl = (*Float)(nil)
	_ SchemaImpl = (*String)(nil)
	_ SchemaImpl = (*Array)(nil)
	_ SchemaImpl = (*Map)(nil)
	_ SchemaImpl = (*Struct)(nil)
	_ Schema     = (*SchemaRef)(nil)
)
