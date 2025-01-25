package golang

import (
	"fmt"
	"go/token"
	"reflect"
	"strings"

	code "github.com/trwk76/gocode"
)

func IsID(s string) bool {
	return token.IsIdentifier(s)
}

func IsKeyword(s string) bool {
	return token.IsKeyword(s)
}

func SymbolFor[T any](unit *Unit) Symbol {
	t := reflect.TypeFor[T]()

	pkg := unit.Imports.Ensure("", t.PkgPath())
	id := t.Name()

	if idx := strings.IndexByte(id, '['); idx > 0 {
		id = id[:idx]
	}

	return Symbol{
		Package: &pkg,
		ID:      ID(id),
	}
}

type (
	Comment string
	PkgName string
	ID      string

	PkgRef struct {
		alias PkgName
		path  string
	}

	Symbol struct {
		Package *PkgRef
		ID      ID
		GenArgs GenArgs
	}

	GenArgs []Type
)

var (
	Ignore ID = ID("_")
)

func (c Comment) write(w *code.Writer) {
	if c == "" {
		return
	}

	for _, line := range strings.Split(string(c), "\n") {
		w.WriteString("//")
		w.WriteString(line)
		w.Newline()
	}
}

func (p PkgName) check() {
	if !token.IsIdentifier(string(p)) || string(p) != strings.ToLower(string(p)) {
		panic(fmt.Errorf("'%s' is not a valid package name", p))
	}
}

func (p PkgName) write(w *code.Writer) {
	p.check()
	w.WriteString(string(p))
}

func (i ID) write(w *code.Writer) {
	if !token.IsIdentifier(string(i)) {
		panic(fmt.Errorf("'%s' is not a valid identifier", i))
	}

	w.WriteString(string(i))
}

func (p PkgRef) write(w *code.Writer) {
	p.alias.check()
	fmt.Fprintf(w, "%s %q", p.alias, p.path)
}

func (s Symbol) simple() bool {
	return s.GenArgs.simpleGenArgs()
}

func (s Symbol) write(w *code.Writer) {
	if s.Package != nil {
		s.Package.alias.check()

		if !token.IsExported(string(s.ID)) {
			panic(fmt.Errorf("unexported symbol '%s' while referencing package '%s' (%s)", s.ID, s.Package.alias, s.Package.path))
		}

		w.WriteString(string(s.Package.alias))
		w.WriteByte('.')
	}

	s.ID.write(w)
	s.GenArgs.write(w)
}

func (a GenArgs) write(w *code.Writer) {
	if len(a) < 1 {
		return
	}

	w.WriteByte('[')

	for idx, itm := range a {
		if idx > 0 {
			w.WriteString(", ")
		}

		itm.writeType(w)
	}

	w.WriteByte(']')
}

func idString(i ID) string {
	return writeString(func(w *code.Writer) { i.write(w) })
}

func writeString(f func(w *code.Writer)) string {
	return code.WriteString("\t", f)
}
