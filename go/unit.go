package golang

import (
	"fmt"
	"strings"

	code "github.com/trwk76/go-code"
)

type (
	Unit struct {
		Prefix  Comment
		Package PkgName
		Imports Imports
		Decls   Decls
	}

	Imports struct {
		sys []PkgRef
		ext []PkgRef
	}
)

func (u Unit) Write(w *code.Writer) {
	u.Prefix.write(w)

	if len(u.Prefix) > 0 {
		w.Newline()
	}

	fmt.Fprintf(w, "package %s", u.Package)
	w.Newline()
	u.Imports.write(w)

	for _, decl := range u.Decls {
		decl.writeDecl(w)
	}

	w.Newline()
}

func (i *Imports) Ensure(alias PkgName, path string) PkgRef {
	if path == "" {
		panic(fmt.Errorf("import path must not be empty"))
	}

	dest := &i.sys
	if !isSysImport(path) {
		dest = &i.ext
	}

	if alias == "" {
		idx := strings.LastIndexByte(path, '/')

		if idx > 0 {
			alias = PkgName(path[idx+1:])
		} else {
			alias = PkgName(path)
		}
	}

	alias.check()

	for idx, imp := range *dest {
		if imp.path == path {
			if imp.alias == PkgName(Ignore) {
				(*dest)[idx].alias = alias
				return imp
			} else if imp.alias == alias {
				return imp
			} else {
				panic(fmt.Errorf("package alias '%s' already exists for package '%s'", alias, path))
			}
		}
	}

	ref := PkgRef{alias: alias, path: path}
	*dest = append(*dest, ref)
	return ref
}

func (i Imports) write(w *code.Writer) {
	total := len(i.sys) + len(i.ext)

	switch total {
	case 0:
		// do nothing
		return
	case 1:
		// import _ "embed"
		var item PkgRef

		if len(i.sys) > 0 {
			item = i.sys[0]
		} else {
			item = i.ext[0]
		}

		w.Newline()
		w.WriteString("import ")
		item.write(w)
		w.Newline()
	default:
		// import ( )
		w.Newline()
		w.WriteString("import (")
		w.Newline()
		w.Indent(func(w *code.Writer) {
			for _, item := range i.sys {
				item.write(w)
				w.Newline()
			}

			if len(i.sys) > 0 && len(i.ext) > 0 {
				w.Newline()
			}

			for _, item := range i.ext {
				item.write(w)
				w.Newline()
			}
		})
		w.WriteByte(')')
		w.Newline()
	}
}

func isSysImport(path string) bool {
	return !strings.Contains(path, ".")
}
