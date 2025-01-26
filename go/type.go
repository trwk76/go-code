package golang

import (
	"errors"
	"strconv"
	"strings"

	code "github.com/trwk76/go-code"
)

var (
	Any        Symbol = Symbol{ID: "any"}
	Bool       Symbol = Symbol{ID: "bool"}
	Byte       Symbol = Symbol{ID: "byte"}
	Comparable Symbol = Symbol{ID: "comparable"}
	Complex64  Symbol = Symbol{ID: "complex64"}
	Complex128 Symbol = Symbol{ID: "complex128"}
	Error      Symbol = Symbol{ID: "error"}
	Float32    Symbol = Symbol{ID: "float32"}
	Float64    Symbol = Symbol{ID: "float64"}
	Int        Symbol = Symbol{ID: "int"}
	Int8       Symbol = Symbol{ID: "int8"}
	Int16      Symbol = Symbol{ID: "int16"}
	Int32      Symbol = Symbol{ID: "int32"}
	Int64      Symbol = Symbol{ID: "int64"}
	Rune       Symbol = Symbol{ID: "rune"}
	String     Symbol = Symbol{ID: "string"}
	Uint       Symbol = Symbol{ID: "uint"}
	Uint8      Symbol = Symbol{ID: "uint8"}
	Uint16     Symbol = Symbol{ID: "uint16"}
	Uint32     Symbol = Symbol{ID: "uint32"}
	Uint64     Symbol = Symbol{ID: "uint64"}
	UintPtr    Symbol = Symbol{ID: "uintptr"}
)

type (
	Type interface {
		TypeSpec
		simpleType() bool
		writeType(w *code.Writer)
	}

	PtrType struct {
		Item Type
	}

	SliceType struct {
		Items Type
		Size  Expr
	}

	MapType struct {
		Key   Type
		Value Type
	}

	InterfaceType struct {
		Consts GenConsts
		Meths  []InterfaceMeth
	}

	StructType struct {
		Bases  []Type
		Fields []StructField
	}

	GenConsts []GenConst

	InterfaceMeth struct {
		ID     ID
		Params Params
		Return Params
	}

	StructField struct {
		Comment Comment
		ID      ID
		Type    Type
		Tags    Tags
	}

	Tag struct {
		Name  string
		Value string
	}

	Tags []Tag
)

func (s Symbol) simpleType() bool {
	return s.simple()
}

func (s Symbol) simpleTypeSpec() bool {
	return s.simpleType()
}

func (s Symbol) writeType(w *code.Writer) {
	s.write(w)
}

func (s Symbol) writeTypeSpec(w *code.Writer) {
	s.write(w)
}

func (t PtrType) simpleType() bool {
	return (t.Item == nil || t.Item.simpleType())
}

func (t PtrType) simpleTypeSpec() bool {
	return t.simpleType()
}

func (t PtrType) writeType(w *code.Writer) {
	w.WriteByte('*')
	writeType(w, t.Item, "pointer type requires an item type")
}

func (t PtrType) writeTypeSpec(w *code.Writer) {
	t.writeType(w)
}

func (t SliceType) simpleType() bool {
	return (t.Size == nil || t.Size.simpleExpr()) &&
		(t.Items == nil || t.Items.simpleType())
}

func (t SliceType) simpleTypeSpec() bool {
	return t.simpleType()
}

func (t SliceType) writeType(w *code.Writer) {
	w.WriteByte('[')
	writeExpr(w, t.Size, true, "")
	w.WriteByte(']')
	writeType(w, t.Items, "slice type requires an item type")
}

func (t SliceType) writeTypeSpec(w *code.Writer) {
	t.writeType(w)
}

func (t MapType) simpleType() bool {
	return (t.Key == nil || t.Key.simpleType()) &&
		(t.Value == nil || t.Value.simpleType())
}

func (t MapType) simpleTypeSpec() bool {
	return t.simpleType()
}

func (s MapType) writeType(w *code.Writer) {
	w.WriteString("map[")
	writeType(w, s.Key, "map type requires a key type")
	w.WriteByte(']')
	writeType(w, s.Value, "map type requires a value type")
}

func (s MapType) writeTypeSpec(w *code.Writer) {
	s.writeType(w)
}

func (t InterfaceType) simpleType() bool {
	return len(t.Consts) < 1 && len(t.Meths) < 1
}

func (t InterfaceType) simpleTypeSpec() bool {
	return t.simpleType()
}

func (t InterfaceType) writeType(w *code.Writer) {
	w.WriteString("interface {")

	if len(t.Consts) > 0 || len(t.Meths) > 0 {
		w.Newline()
		w.Indent(func(w *code.Writer) {
			if len(t.Consts) > 0 {
				for idx, itm := range t.Consts {
					if idx > 0 {
						w.WriteString(" | ")
					}

					itm.writeConstraint(w)
				}

				w.Newline()

				if len(t.Meths) > 0 {
					w.Newline()
				}
			}

			for _, itm := range t.Meths {
				itm.write(w)
				w.Newline()
			}
		})
	}

	w.WriteByte('}')
}

func (t InterfaceType) writeTypeSpec(w *code.Writer) {
	t.writeType(w)
}

func (m InterfaceMeth) write(w *code.Writer) {
	m.ID.write(w)
	m.Params.write(w, true)
	m.Return.write(w, false)
}

func (t StructType) simpleType() bool {
	return len(t.Bases) < 1 && len(t.Fields) < 1
}

func (t StructType) simpleTypeSpec() bool {
	return t.simpleType()
}

func (t StructType) writeType(w *code.Writer) {
	w.WriteString("struct {")

	if len(t.Bases) > 0 || len(t.Fields) > 0 {
		w.Newline()
		w.Indent(func(w *code.Writer) {
			rows := make([]code.TableRow, 0, len(t.Bases)+len(t.Fields)+1)

			for _, base := range t.Bases {
				rows = append(rows, code.TableRow{
					Columns: []string{typeString(base, "base must not be null")},
				})
			}

			if len(t.Bases) > 0 && len(t.Fields) > 0 {
				rows = append(rows, code.TableRow{})
			}

			for _, fld := range t.Fields {
				cols := []string{
					idString(fld.ID),
					typeString(fld.Type, "struct field requires a type"),
				}

				if tag := fld.Tags.String(); tag != "" {
					cols = append(cols, tag)
				}

				rows = append(rows, code.TableRow{
					Prefix:  string(fld.Comment),
					Columns: cols,
				})
			}

			w.Table(rows...)
		})
	}

	w.WriteByte('}')
}

func (t StructType) writeTypeSpec(w *code.Writer) {
	t.writeType(w)
}

func (t Tag) String() string {
	if t.Name == "" {
		panic(errors.New("tag name must not be empty"))
	}

	return t.Name + ":" + strconv.Quote(t.Value)
}

func (t Tags) String() string {
	if len(t) < 1 {
		return ""
	}

	items := make([]string, len(t))

	for idx, itm := range t {
		items[idx] = itm.String()
	}

	return "`" + strings.Join(items, " ") + "`"
}

func (a GenArgs) simpleGenArgs() bool {
	for _, itm := range a {
		if itm != nil && !itm.simpleType() {
			return false
		}
	}

	return true
}

func typeString(t Type, reqMessage string) string {
	return writeString(func(w *code.Writer) { writeType(w, t, reqMessage) })
}

func writeType(w *code.Writer, t Type, reqMessage string) {
	if t == nil {
		if reqMessage != "" {
			panic(errors.New(reqMessage))
		}

		return
	}

	t.writeType(w)
}

var (
	_ Type = Symbol{}
	_ Type = PtrType{}
	_ Type = SliceType{}
	_ Type = MapType{}
	_ Type = InterfaceType{}
	_ Type = StructType{}
)
