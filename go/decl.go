package golang

import (
	"fmt"

	code "github.com/trwk76/gocode"
)

type (
	ConstDecl struct {
		Comment Comment
		ID      ID
		Type    Type
		Value   Expr
	}

	FuncDecl struct {
		Comment   Comment
		ID        ID
		GenParams GenParams
		Params    Params
		Return    Params
		Body      BlockStmt
	}

	MethDecl struct {
		Comment  Comment
		Receiver Param
		ID       ID
		Params   Params
		Return   Params
		Body     BlockStmt
	}

	TypeDecl struct {
		Comment   Comment
		ID        ID
		GenParams GenParams
		Spec      TypeSpec
	}

	VarDecl struct {
		Comment Comment
		ID      ID
		Type    Type
		Value   Expr
	}

	Decl interface {
		writeDecl(w *code.Writer)
	}

	Decls []Decl

	TypeSpec interface {
		simpleTypeSpec() bool
		writeTypeSpec(w *code.Writer)
	}

	TypeAlias struct {
		Target Type
	}

	ConstDecls []ConstDecl
	FuncDecls  []FuncDecl
	MethDecls  []MethDecl
	TypeDecls  []TypeDecl
	VarDecls   []VarDecl

	GenParam struct {
		ID    ID
		Const GenConst
	}

	GenParams []GenParam

	GenConst struct {
		Tilde bool
		Base  Type
	}

	Param struct {
		ID   ID
		Type Type
	}

	Params []Param

	declItem interface {
		simpleDeclItem() bool
		declItemRow() code.TableRow
		writeDeclItem(w *code.Writer, keyword bool)
	}
)

func (d ConstDecl) simpleDeclItem() bool {
	res := true

	if d.Type != nil {
		res = res && d.Type.simpleType()
	}

	if d.Value != nil {
		res = res && d.Value.simpleExpr()
	}

	return res
}

func (d ConstDecl) declItemRow() code.TableRow {
	res := code.TableRow{
		Prefix: string(d.Comment),
		Columns: []string{
			idString(d.ID),
			typeString(d.Type, "constant declaration requires a type"),
		},
	}

	if d.Value != nil {
		res.Columns = append(res.Columns, "= "+exprString(d.Value, ""))
	}

	return res
}

func (d ConstDecl) writeDeclItem(w *code.Writer, keyword bool) {
	d.Comment.write(w)

	if keyword {
		w.WriteString("const ")
	}

	d.ID.write(w)
	w.Space()
	d.Type.writeType(w)

	if d.Value != nil {
		w.WriteString(" = ")
		d.Value.writeExpr(w, false)
	}
}

func (d FuncDecl) simpleDeclItem() bool {
	return d.GenParams.simpleGenParams() &&
		d.Params.simpleParams() &&
		d.Return.simpleParams() &&
		d.Body.simpleStmt()
}

func (d FuncDecl) declItemRow() code.TableRow {
	return code.TableRow{
		Prefix: string(d.Comment),
		Columns: []string{
			"func",
			idString(d.ID) + genParamsString(d.GenParams) + paramsString(d.Params, true) + paramsString(d.Return, false),
			stmtString(d.Body, true, "function declaration requires a body"),
		},
	}
}

func (d FuncDecl) writeDeclItem(w *code.Writer, keyword bool) {
	d.Comment.write(w)

	w.WriteString("func ")
	d.ID.write(w)
	d.GenParams.write(w)
	d.Params.write(w, true)
	d.Return.write(w, false)
	w.Space()
	writeStmt(w, d.Body, false, "function declaration requires a body")
}

func (d MethDecl) simpleDeclItem() bool {
	return d.Receiver.simpleParam() &&
		d.Params.simpleParams() &&
		d.Return.simpleParams() &&
		d.Body.simpleStmt()
}

func (d MethDecl) declItemRow() code.TableRow {
	return code.TableRow{
		Prefix: string(d.Comment),
		Columns: []string{
			"func",
			paramsString(Params{d.Receiver}, true),
			idString(d.ID) + paramsString(d.Params, true) + paramsString(d.Return, false),
			stmtString(d.Body, true, "method declaration requires a body"),
		},
	}
}

func (d MethDecl) writeDeclItem(w *code.Writer, keyword bool) {
	d.Comment.write(w)

	w.WriteString("func ")
	Params{d.Receiver}.write(w, true)
	d.ID.write(w)
	d.Params.write(w, true)
	d.Return.write(w, false)
	w.Space()
	writeStmt(w, d.Body, false, "method declaration requires a body")
}

func (d TypeDecl) simpleDeclItem() bool {
	return d.GenParams.simpleGenParams() && d.Spec.simpleTypeSpec()
}

func (d TypeDecl) declItemRow() code.TableRow {
	return code.TableRow{
		Prefix: string(d.Comment),
		Columns: []string{
			idString(d.ID) + genParamsString(d.GenParams),
			typeSpecString(d.Spec),
		},
	}
}

func (d TypeDecl) writeDeclItem(w *code.Writer, keyword bool) {
	d.Comment.write(w)

	if keyword {
		w.WriteString("type ")
	}

	d.ID.write(w)
	d.GenParams.write(w)
	w.Space()
	writeTypeSpec(w, d.Spec)
}

func (d VarDecl) simpleDeclItem() bool {
	res := true

	if d.Type != nil {
		res = res && d.Type.simpleType()
	}

	if d.Value != nil {
		res = res && d.Value.simpleExpr()
	}

	return res
}

func (d VarDecl) declItemRow() code.TableRow {
	res := code.TableRow{
		Prefix: string(d.Comment),
		Columns: []string{
			idString(d.ID),
			typeString(d.Type, "variable declaration requires a type"),
		},
	}

	if d.Value != nil {
		res.Columns = append(res.Columns, "= "+exprString(d.Value, ""))
	}

	return res
}

func (d VarDecl) writeDeclItem(w *code.Writer, keyword bool) {
	d.Comment.write(w)

	if keyword {
		w.WriteString("var ")
	}

	d.ID.write(w)
	w.Space()
	d.Type.writeType(w)

	if d.Value != nil {
		w.WriteString(" = ")
		d.Value.writeExpr(w, false)
	}
}

func (a TypeAlias) simpleTypeSpec() bool {
	return a.Target.simpleType()
}

func (a TypeAlias) writeTypeSpec(w *code.Writer) {
	w.WriteString("= ")
	a.Target.writeType(w)
}

func typeSpecString(s TypeSpec) string {
	if s == nil {
		panic(fmt.Errorf("type specifier missing"))
	}

	return writeString(func(w *code.Writer) { s.writeTypeSpec(w) })
}

func writeTypeSpec(w *code.Writer, s TypeSpec) {
	if s == nil {
		panic(fmt.Errorf("type specifier missing"))
	}

	s.writeTypeSpec(w)
}

func (c Comment) writeDecl(w *code.Writer) {
	c.write(w)
}

func (d ConstDecls) writeDecl(w *code.Writer) {
	writeDeclItemSection(w, d, "const")
}

func (d FuncDecls) writeDecl(w *code.Writer) {
	writeDeclItems(w, d)
}

func (d MethDecls) writeDecl(w *code.Writer) {
	writeDeclItems(w, d)
}

func (d TypeDecls) writeDecl(w *code.Writer) {
	writeDeclItemSection(w, d, "type")
}

func (d VarDecls) writeDecl(w *code.Writer) {
	writeDeclItemSection(w, d, "var")
}

func (p GenParam) simpleGenParam() bool {
	return p.Const.simpleConstraint()
}

func (p GenParam) write(w *code.Writer) {
	p.ID.write(w)
	w.Space()
	p.Const.writeConstraint(w)
}

func (p GenParams) simpleGenParams() bool {
	for _, itm := range p {
		if !itm.simpleGenParam() {
			return false
		}
	}

	return true
}

func (p GenParams) write(w *code.Writer) {
	if len(p) < 1 {
		return
	}

	w.WriteByte('[')

	for idx, itm := range p {
		if idx > 0 {
			w.WriteString(", ")
		}

		itm.write(w)
	}

	w.WriteByte(']')
}

func genParamsString(p GenParams) string {
	return writeString(func(w *code.Writer) { p.write(w) })
}

func (c GenConst) simpleConstraint() bool {
	return c.Base == nil || c.Base.simpleType()
}

func (c GenConst) writeConstraint(w *code.Writer) {
	if c.Tilde {
		w.WriteByte('~')
	}

	writeType(w, c.Base, "generic contraint requires a base type")
}

func (p Param) simpleParam() bool {
	return p.Type == nil || p.Type.simpleType()
}

func (p Param) write(w *code.Writer) {
	reqMsg := "unnamed parameter requires a type"

	if p.ID != "" {
		p.ID.write(w)
		reqMsg = ""

		if p.Type != nil {
			w.Space()
		}
	}

	writeType(w, p.Type, reqMsg)
}

func (p Params) simpleParams() bool {
	for _, itm := range p {
		if !itm.simpleParam() {
			return false
		}
	}

	return true
}

func (p Params) write(w *code.Writer, forceParens bool) {
	if !forceParens {
		w.Space()
		forceParens = len(p) > 0 && p[0].ID != ""
	}

	if forceParens {
		w.WriteByte('(')
	}

	for idx, itm := range p {
		if idx > 0 {
			w.WriteString(", ")
		}

		itm.write(w)
	}

	if forceParens {
		w.WriteByte(')')
	}
}

func paramsString(p Params, forceParens bool) string {
	return writeString(func(w *code.Writer) { p.write(w, forceParens) })
}

func writeDeclItemSection[T declItem](w *code.Writer, items []T, keyword string) {
	switch len(items) {
	case 0:
		return
	case 1:
		w.Newline()
		items[0].writeDeclItem(w, true)
		w.Newline()
	default:
		w.Newline()
		w.WriteString(keyword + " (")
		w.Newline()
		w.Indent(func(w *code.Writer) {
			writeDeclItems(w, items)
		})
		w.WriteByte(')')
		w.Newline()
	}
}

func writeDeclItems[T declItem](w *code.Writer, items []T) {
	first := true

	for len(items) > 0 {
		if first {
			first = false
		} else {
			w.Newline()
		}

		if cnt := countSimpleDeclItems(items); cnt > 0 {
			// Make table out of these items
			rows := make([]code.TableRow, cnt)

			for idx, itm := range items[:cnt] {
				rows[idx] = itm.declItemRow()
			}

			w.Table(rows...)

			items = items[cnt:]
		} else {
			// Write item
			item := items[0]

			item.writeDeclItem(w, false)
			w.Newline()

			items = items[1:]
		}
	}
}

func countSimpleDeclItems[T declItem](items []T) int {
	for idx, itm := range items {
		if !itm.simpleDeclItem() {
			return idx
		}
	}

	return len(items)
}

var (
	_ declItem = ConstDecl{}
	_ declItem = FuncDecl{}
	_ declItem = MethDecl{}
	_ declItem = TypeDecl{}
	_ declItem = VarDecl{}
	_ TypeSpec = TypeAlias{}
	_ Decl     = Comment("")
	_ Decl     = ConstDecls{}
	_ Decl     = FuncDecls{}
	_ Decl     = MethDecls{}
	_ Decl     = TypeDecls{}
	_ Decl     = VarDecls{}
)
