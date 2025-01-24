package golang

import (
	"errors"

	code "github.com/trwk76/gocode"
)

type (
	Stmt interface {
		simpleStmt() bool
		writeStmt(w *code.Writer, singleLine bool)
	}

	ElseStmt interface {
		Stmt
		elseStmt()
	}

	InitStmt interface {
		Stmt
		initStmt()
	}

	AssignStmt struct {
		Auto  bool
		Dests Exprs
		Srcs  Exprs
	}

	BlockStmt []Stmt

	BreakStmt    struct{}
	ContinueStmt struct{}

	DeferStmt struct {
		Expr Expr
	}

	ExprStmt struct {
		Expr Expr
	}

	FallThroughStmt struct{}

	ForStmt struct {
		Init InitStmt
		Cond Expr
		Next *AssignStmt
		Then BlockStmt
	}

	IfStmt struct {
		Init InitStmt
		Cond Expr
		Then BlockStmt
		Else ElseStmt
	}

	ReturnStmt struct {
		Value Expr
	}

	SwitchStmt struct {
		Value Expr
		Cases []SwitchCase
	}

	SwitchCase struct {
		Value Expr
		Stmts []Stmt
	}
)

func (s AssignStmt) initStmt() {}

func (s AssignStmt) simpleStmt() bool {
	return s.Dests.simpleExprs() && s.Srcs.simpleExprs()
}

func (s AssignStmt) writeStmt(w *code.Writer, singleLine bool) {
	s.Dests.writeExprs(w, true)
	w.Space()
	if s.Auto {
		w.WriteString(":=")
	} else {
		w.WriteString("=")
	}
	w.Space()
	s.Srcs.writeExprs(w, singleLine)
}

func (s BlockStmt) elseStmt() {}

func (s BlockStmt) simpleStmt() bool {
	for _, itm := range s {
		if itm != nil && !itm.simpleStmt() {
			return false
		}
	}

	return len(s) < 3
}

func (s BlockStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteByte('{')

	if len(s) > 0 {
		if singleLine {
			w.Space()

			first := true

			for _, itm := range s {
				if itm != nil {
					if first {
						first = false
					} else {
						w.WriteString("; ")
					}

					writeStmt(w, itm, true, "")
				}
			}

			if !first {
				w.Space()
			}
		} else {
			w.Newline()

			w.Indent(func(w *code.Writer) {
				for _, itm := range s {
					writeStmt(w, itm, false, "")
					w.Newline()
				}
			})
		}
	}

	w.WriteByte('}')
}

func (BreakStmt) simpleStmt() bool {
	return true
}

func (BreakStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("break")
}

func (ContinueStmt) simpleStmt() bool {
	return true
}

func (ContinueStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("continue")
}

func (s DeferStmt) simpleStmt() bool {
	return s.Expr == nil || s.Expr.simpleExpr()
}

func (s DeferStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("defer ")
	writeExpr(w, s.Expr, singleLine, "defer statement requires an expression")
}

func (s ExprStmt) simpleStmt() bool {
	return s.Expr == nil || s.Expr.simpleExpr()
}

func (s ExprStmt) writeStmt(w *code.Writer, singleLine bool) {
	writeExpr(w, s.Expr, singleLine, "expression statement requires an expression")
}

func (FallThroughStmt) simpleStmt() bool {
	return true
}

func (FallThroughStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("fallthrough")
}

func (s ForStmt) simpleStmt() bool {
	return (s.Init == nil || s.Init.simpleStmt()) &&
		(s.Cond == nil || s.Cond.simpleExpr()) &&
		(s.Next == nil || s.Next.simpleStmt()) &&
		s.Then.simpleStmt()
}

func (s ForStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("for ")

	if s.Init != nil || s.Next != nil {
		writeStmt(w, s.Init, singleLine, "")
		w.WriteString("; ")
		writeExpr(w, s.Cond, singleLine, "")
		w.WriteString("; ")

		if s.Next != nil {
			writeStmt(w, s.Next, singleLine, "")
			w.Space()
		}
	} else {
		if s.Cond != nil {
			writeExpr(w, s.Cond, singleLine, "")
			w.Space()
		}
	}

	s.Then.writeStmt(w, singleLine)
}

func (s IfStmt) elseStmt() {}

func (s IfStmt) simpleStmt() bool {
	return (s.Cond == nil || s.Cond.simpleExpr()) ||
		s.Then.simpleStmt() ||
		(s.Else == nil || s.Else.simpleStmt())
}

func (s IfStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("if ")

	if s.Init != nil {
		s.Init.writeStmt(w, true)
		w.WriteString("; ")
	}

	writeExpr(w, s.Cond, true, "if statement requires a condition")
	w.Space()

	s.Then.writeStmt(w, singleLine)

	if s.Else != nil {
		w.WriteString(" else ")
		writeStmt(w, s.Else, singleLine, "")
	}
}

func (s ReturnStmt) simpleStmt() bool {
	return s.Value == nil || s.Value.simpleExpr()
}

func (s ReturnStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("return")

	if s.Value != nil {
		w.Space()
		writeExpr(w, s.Value, singleLine, "")
	}
}

func (s SwitchStmt) simpleStmt() bool {
	return false
}

func (s SwitchStmt) writeStmt(w *code.Writer, singleLine bool) {
	w.WriteString("switch ")
	writeExpr(w, s.Value, singleLine, "switch statement requires an expression")
	w.WriteString(" {")

	if singleLine {
		for _, cas := range s.Cases {
			if cas.Value != nil {
				w.WriteString(" case ")
				writeExpr(w, cas.Value, singleLine, "")
				w.WriteString(": ")
			} else {
				w.WriteString("default: ")
			}

			w.Indent(func(w *code.Writer) {
				for _, stmt := range cas.Stmts {
					writeStmt(w, stmt, singleLine, "")
					w.WriteString("; ")
				}
			})
		}
	} else {
		w.Newline()

		for _, cas := range s.Cases {
			if cas.Value != nil {
				w.WriteString("case ")
				writeExpr(w, cas.Value, singleLine, "")
				w.WriteByte(':')
			} else {
				w.WriteString("default:")
			}

			w.Indent(func(w *code.Writer) {
				for _, stmt := range cas.Stmts {
					w.Newline()
					writeStmt(w, stmt, singleLine, "")
				}
			})
		}
	}

	w.WriteByte('}')
}

func (d ConstDecl) simpleStmt() bool {
	return d.simpleDeclItem()
}

func (d ConstDecl) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDeclItem(w, true)
}

func (d ConstDecls) simpleStmt() bool {
	return len(d) < 2
}

func (d ConstDecls) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDecl(w)
}

func (d FuncDecl) simpleStmt() bool {
	return d.simpleDeclItem()
}

func (d FuncDecl) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDeclItem(w, true)
}

func (d TypeDecl) simpleStmt() bool {
	return d.simpleDeclItem()
}

func (d TypeDecl) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDeclItem(w, true)
}

func (d TypeDecls) simpleStmt() bool {
	return len(d) < 2
}

func (d TypeDecls) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDecl(w)
}

func (d VarDecl) initStmt() {}

func (d VarDecl) simpleStmt() bool {
	return d.simpleDeclItem()
}

func (d VarDecl) writeStmt(w *code.Writer, singleLine bool) {
	d.Comment.write(w)

	w.WriteString("var ")
	d.ID.write(w)
	w.Space()
	d.Type.writeType(w)

	if d.Value != nil {
		w.WriteString(" = ")
		d.Value.writeExpr(w, singleLine)
	}
}

func (d VarDecls) simpleStmt() bool {
	return len(d) < 2
}

func (d VarDecls) writeStmt(w *code.Writer, singleLine bool) {
	d.writeDecl(w)
}

func stmtString(s Stmt, singleLine bool, reqMessage string) string {
	return writeString(func(w *code.Writer) { writeStmt(w, s, singleLine, reqMessage) })
}

func writeStmt(w *code.Writer, s Stmt, singleLine bool, reqMessage string) {
	if s == nil {
		if reqMessage != "" {
			panic(errors.New(reqMessage))
		}

		return
	}

	s.writeStmt(w, singleLine)
}

var (
	_ InitStmt = AssignStmt{}
	_ ElseStmt = BlockStmt{}
	_ Stmt     = BreakStmt{}
	_ Stmt     = ContinueStmt{}
	_ Stmt     = DeferStmt{}
	_ Stmt     = ExprStmt{}
	_ Stmt     = FallThroughStmt{}
	_ Stmt     = ForStmt{}
	_ ElseStmt = IfStmt{}
	_ Stmt     = ReturnStmt{}
	_ Stmt     = SwitchStmt{}
	_ Stmt     = ConstDecl{}
	_ Stmt     = ConstDecls{}
	_ Stmt     = FuncDecl{}
	_ Stmt     = TypeDecl{}
	_ Stmt     = TypeDecls{}
	_ InitStmt = VarDecl{}
	_ Stmt     = VarDecls{}
)
