package stdhttp

import (
	"fmt"
	"math"
	"reflect"

	code "github.com/trwk76/go-code"
	g "github.com/trwk76/go-code/go"
	"github.com/trwk76/go-code/web/api"
)

type (
	TypeConverter interface {
		api.SchemaImplVisitor

		init(typ reflect.Type, name string)
		result() g.Type
	}

	DefaultTypeConverter struct {
		typ  reflect.Type
		name string
		res  g.Type
	}
)

func (c *DefaultTypeConverter) Boolean(i *api.Boolean, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	return g.Bool
}

func (c *DefaultTypeConverter) Enum(i *api.Enum, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	switch i.Type.Kind() {
	case reflect.Bool:
		return g.Bool
	case reflect.Float32:
		return g.Float32
	case reflect.Float64:
		return g.Float64
	case reflect.Int16:
		return g.Int16
	case reflect.Int32:
		return g.Int32
	case reflect.Int, reflect.Int64:
		return g.Int64
	case reflect.Int8:
		return g.Int8
	case reflect.String:
		return g.String
	case reflect.Uint16:
		return g.Uint16
	case reflect.Uint32:
		return g.Uint32
	case reflect.Uint, reflect.Uint64:
		return g.Uint64
	case reflect.Uint8:
		return g.Uint8
	}

	panic(fmt.Errorf("enumeration type '%s' is not supported", i.Type.String()))
}

func (c *DefaultTypeConverter) Integer(i *api.Integer, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	if i.Minimum != 0 || i.Maximum != 0 {
		if i.Minimum >= 0 && i.Maximum <= math.MaxUint8 {
			return g.Uint8
		} else if i.Minimum >= math.MinInt8 && i.Maximum <= math.MaxInt8 {
			return g.Int8
		} else if i.Minimum >= 0 && i.Maximum <= math.MaxUint16 {
			return g.Uint16
		} else if i.Minimum >= math.MinInt16 && i.Maximum <= math.MaxInt16 {
			return g.Int16
		} else if i.Minimum >= 0 && i.Maximum <= math.MaxUint32 {
			return g.Uint32
		} else if i.Minimum >= math.MinInt32 && i.Maximum <= math.MaxInt32 {
			return g.Int32
		}
	}

	return g.Int64
}

func (c *DefaultTypeConverter) Uinteger(i *api.Uinteger, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	if i.Minimum != 0 || i.Maximum != 0 {
		if i.Minimum >= 0 && i.Maximum <= math.MaxUint8 {
			return g.Uint8
		} else if i.Minimum >= 0 && i.Maximum <= math.MaxUint16 {
			return g.Uint16
		} else if i.Minimum >= 0 && i.Maximum <= math.MaxUint32 {
			return g.Uint32
		}
	}

	return g.Uint64
}

func (c *DefaultTypeConverter) Float(i *api.Float, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	if i.Minimum != 0 || i.Maximum != 0 {
		if i.Minimum >= -math.MaxFloat32 && i.Maximum <= math.MaxFloat32 {
			return g.Float32
		}
	}

	return g.Float64
}

func (c *DefaultTypeConverter) String(i *api.String, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	return g.String
}

func (c *DefaultTypeConverter) Array(i *api.Array, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	return g.SliceType{Items: c.Convert(i.Items)}
}

func (c *DefaultTypeConverter) Map(i *api.Map, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	return g.MapType{
		Key:   c.Convert(i.Key),
		Value: c.Convert(i.Value),
	}
}

func (c *DefaultTypeConverter) Struct(i *api.Struct, name string) g.Type {
	if name != "" {
		return g.Symbol{ID: g.ID(name)}
	}

	bases := make([]g.Type, 0)
	flds := make([]g.StructField, 0)

	for _, base := range i.Bases {
		bases = append(bases, c.Convert(base))
	}

	for _, fld := range i.Fields {
		tag := fld.Name
		typ := c.Convert(fld.Schema)

		if fld.Optional {
			tag += ",omitempty"
			typ = g.PtrType{Item: typ}
		}

		flds = append(flds, g.StructField{
			ID:   g.ID(code.IDToPascal(fld.Name)),
			Type: typ,
			Tags: g.Tags{
				{
					Name:  "json",
					Value: tag,
				},
			},
		})
	}

	return g.StructType{
		Bases:  bases,
		Fields: flds,
	}
}

func (c *DefaultTypeConverter) Convert(i api.Schema) g.Type {
	if ref, ok := i.(*api.SchemaRef); ok {
		return g.Symbol{ID: g.ID(ref.Key())}
	} else if impl, ok := i.(api.SchemaImpl); ok {
		return convertType(c.typ, impl, "")
	}

	panic(fmt.Errorf("api schema type '%T' not supported", i))
}

func convertType(t reflect.Type, i api.SchemaImpl, name string) g.Type {
	cnv := reflect.New(t).Interface().(TypeConverter)

	cnv.init(t, name)
	i.Accept(cnv)

	return cnv.result()
}

func (c *DefaultTypeConverter) init(typ reflect.Type, name string) {
	c.typ = typ
	c.name = name
}

func (c *DefaultTypeConverter) result() g.Type {
	return c.res
}

func (c *DefaultTypeConverter) VisitBoolean(i *api.Boolean)   { c.res = c.Boolean(i, c.name) }
func (c *DefaultTypeConverter) VisitEnum(i *api.Enum)         { c.res = c.Enum(i, c.name) }
func (c *DefaultTypeConverter) VisitInteger(i *api.Integer)   { c.res = c.Integer(i, c.name) }
func (c *DefaultTypeConverter) VisitUinteger(i *api.Uinteger) { c.res = c.Uinteger(i, c.name) }
func (c *DefaultTypeConverter) VisitFloat(i *api.Float)       { c.res = c.Float(i, c.name) }
func (c *DefaultTypeConverter) VisitString(i *api.String)     { c.res = c.String(i, c.name) }
func (c *DefaultTypeConverter) VisitArray(i *api.Array)       { c.res = c.Array(i, c.name) }
func (c *DefaultTypeConverter) VisitMap(i *api.Map)           { c.res = c.Map(i, c.name) }
func (c *DefaultTypeConverter) VisitStruct(i *api.Struct)     { c.res = c.Struct(i, c.name) }

var (
	_ TypeConverter = (*DefaultTypeConverter)(nil)
)
