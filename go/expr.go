package golang

import (
	"errors"
	"strconv"

	code "github.com/trwk76/go-code"
)

var (
	Nil   NilExpr
	Iota  IotaExpr
	False BoolExpr = BoolExpr(false)
	True  BoolExpr = BoolExpr(true)
)

type (
	Expr interface {
		simpleExpr() bool
		writeExpr(w *code.Writer, singleLine bool)
	}

	NilExpr    struct{}
	IotaExpr   struct{}
	BoolExpr   bool
	IntExpr    int64
	UintExpr   uint64
	FloatExpr  float64
	RuneExpr   rune
	StringExpr string

	ParExpr struct {
		Expr Expr
	}

	CastExpr struct {
		Type  Type
		Value Expr
	}

	SliceExpr struct {
		Type  Type
		Items Exprs
	}

	MapExpr struct {
		Type    Type
		Entries []MapEntry
	}

	MapEntry struct {
		Key   Expr
		Value Expr
	}

	StructExpr struct {
		Type   Type
		Fields []StructExprField
	}

	StructExprField struct {
		ID    ID
		Value Expr
	}

	FuncExpr struct {
		Params Params
		Return Params
		Body   BlockStmt
	}

	NewExpr struct {
		Type Type
	}

	MakeExpr struct {
		Type  Type
		Sizes Exprs
	}

	MemberExpr struct {
		Value Expr
		ID    ID
	}

	CallExpr struct {
		Func Expr
		Args Exprs
	}

	IndexExpr struct {
		Slice Expr
		Index Expr
	}

	RangeExpr struct {
		Slice Expr
		Min   Expr
		Max   Expr
	}

	IdentExpr struct {
		Op Expr
	}

	NegateExpr struct {
		Op Expr
	}

	NotExpr struct {
		Op Expr
	}

	ComplementExpr struct {
		Op Expr
	}

	AddrOfExpr struct {
		Op Expr
	}

	DerefExpr struct {
		Op Expr
	}

	AddExpr struct {
		LHS Expr
		RHS Expr
	}

	SubtractExpr struct {
		LHS Expr
		RHS Expr
	}

	MultiplyExpr struct {
		LHS Expr
		RHS Expr
	}

	DivideExpr struct {
		LHS Expr
		RHS Expr
	}

	ModulusExpr struct {
		LHS Expr
		RHS Expr
	}

	ShiftLeftExpr struct {
		LHS Expr
		RHS Expr
	}

	ShiftRightExpr struct {
		LHS Expr
		RHS Expr
	}

	EqualExpr struct {
		LHS Expr
		RHS Expr
	}

	NotEqualExpr struct {
		LHS Expr
		RHS Expr
	}

	LessThanExpr struct {
		LHS Expr
		RHS Expr
	}

	LessOrEqualExpr struct {
		LHS Expr
		RHS Expr
	}

	MoreThanExpr struct {
		LHS Expr
		RHS Expr
	}

	MoreOrEqualExpr struct {
		LHS Expr
		RHS Expr
	}

	BitAndExpr struct {
		LHS Expr
		RHS Expr
	}

	BitXorExpr struct {
		LHS Expr
		RHS Expr
	}

	BitOrExpr struct {
		LHS Expr
		RHS Expr
	}

	LogAndExpr struct {
		LHS Expr
		RHS Expr
	}

	LogOrExpr struct {
		LHS Expr
		RHS Expr
	}

	Exprs []Expr
)

func (s Symbol) simpleExpr() bool {
	return s.simple()
}

func (s Symbol) writeExpr(w *code.Writer, singleLine bool) {
	s.write(w)
}

func (e NilExpr) simpleExpr() bool {
	return true
}

func (e NilExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString("nil")
}

func (e IotaExpr) simpleExpr() bool {
	return true
}

func (e IotaExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString("iota")
}

func (e BoolExpr) simpleExpr() bool {
	return true
}

func (e BoolExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.FormatBool(bool(e)))
}

func (e IntExpr) simpleExpr() bool {
	return true
}

func (e IntExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.FormatInt(int64(e), 10))
}

func (e UintExpr) simpleExpr() bool {
	return true
}

func (e UintExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.FormatUint(uint64(e), 10))
}

func (e FloatExpr) simpleExpr() bool {
	return true
}

func (e FloatExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.FormatFloat(float64(e), 'g', -1, 64))
}

func (e RuneExpr) simpleExpr() bool {
	return true
}

func (e RuneExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.QuoteRune(rune(e)))
}

func (e StringExpr) simpleExpr() bool {
	return true
}

func (e StringExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString(strconv.Quote(string(e)))
}

func (e ParExpr) simpleExpr() bool {
	return e.Expr == nil || e.Expr.simpleExpr()
}

func (e ParExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('(')
	writeExpr(w, e.Expr, singleLine, "parenthesis expression requires an inner expression")
	w.WriteByte(')')
}

func (e CastExpr) simpleExpr() bool {
	return (e.Type == nil || e.Type.simpleType()) &&
		(e.Value == nil || e.Value.simpleExpr())
}

func (e CastExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('(')
	writeType(w, e.Type, "cast expression requires a target type")
	w.WriteString(")(")
	writeExpr(w, e.Value, singleLine, "cast expression requires a value expression")
	w.WriteByte(')')
}

func (e SliceExpr) simpleExpr() bool {
	return (e.Type == nil || e.Type.simpleType()) &&
		e.Items.simpleExprs()
}

func (e SliceExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeType(w, e.Type, "")
	w.WriteByte('{')
	e.Items.writeExprs(w, singleLine)
	w.WriteByte('}')
}

func (e MapExpr) simpleExpr() bool {
	for _, itm := range e.Entries {
		if itm.Key != nil && !itm.Key.simpleExpr() {
			return false
		}

		if itm.Value != nil && !itm.Value.simpleExpr() {
			return false
		}
	}

	return (e.Type == nil || e.Type.simpleType())
}

func (e MapExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeType(w, e.Type, "")
	w.WriteByte('{')

	if len(e.Entries) > 0 {
		if singleLine {
			for idx, itm := range e.Entries {
				if idx > 0 {
					w.WriteString(", ")
				}

				writeExpr(w, itm.Key, singleLine, "key in map entry must not be nil")
				w.WriteString(": ")
				writeExpr(w, itm.Value, singleLine, "value in map entry must not be nil")
			}
		} else {
			w.Newline()
			w.Indent(func(w *code.Writer) {
				for _, itm := range e.Entries {
					writeExpr(w, itm.Key, singleLine, "key in map entry must not be nil")
					w.WriteString(": ")
					writeExpr(w, itm.Value, singleLine, "value in map entry must not be nil")
					w.WriteByte(',')
					w.Newline()
				}
			})
		}
	}

	w.WriteByte('}')
}

func (e StructExpr) simpleExpr() bool {
	for _, itm := range e.Fields {
		if itm.Value != nil && !itm.Value.simpleExpr() {
			return false
		}
	}

	return (e.Type == nil || e.Type.simpleType())
}

func (e StructExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeType(w, e.Type, "")
	w.WriteByte('{')

	if len(e.Fields) > 0 {
		if singleLine {
			for idx, itm := range e.Fields {
				if idx > 0 {
					w.WriteString(", ")
				}

				itm.ID.write(w)
				w.WriteString(": ")
				writeExpr(w, itm.Value, singleLine, "value in struct field must not be nil")
			}
		} else {
			w.Newline()
			w.Indent(func(w *code.Writer) {
				for _, itm := range e.Fields {
					itm.ID.write(w)
					w.WriteString(": ")
					writeExpr(w, itm.Value, singleLine, "value in struct field must not be nil")
					w.WriteByte(',')
					w.Newline()
				}
			})
		}
	}

	w.WriteByte('}')
}

func (e NewExpr) simpleExpr() bool {
	return (e.Type == nil || e.Type.simpleType())
}

func (e NewExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString("new(")
	writeType(w, e.Type, "new function requires a type")
	w.WriteByte(')')
}

func (e FuncExpr) simpleExpr() bool {
	return e.Params.simpleParams() && e.Return.simpleParams() && e.Body.simpleStmt()
}

func (e FuncExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString("func")
	e.Params.write(w, true)
	e.Return.write(w, false)
	w.Space()
	e.Body.writeStmt(w, singleLine)
}

func (e MakeExpr) simpleExpr() bool {
	for _, itm := range e.Sizes {
		if itm != nil && !itm.simpleExpr() {
			return false
		}
	}

	return (e.Type == nil || e.Type.simpleType())
}

func (e MakeExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteString("make(")
	writeType(w, e.Type, "make function requires a type")

	for _, itm := range e.Sizes {
		w.WriteString(", ")
		writeExpr(w, itm, singleLine, "make length argument must not be nil")
	}

	w.WriteByte(')')
}

func (e MemberExpr) simpleExpr() bool {
	return e.Value == nil || e.Value.simpleExpr()
}

func (e MemberExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.Value, singleLine, "member expression requires a value expression")
	w.WriteByte('.')
	e.ID.write(w)
}

func (e CallExpr) simpleExpr() bool {
	for _, arg := range e.Args {
		if !arg.simpleExpr() {
			return false
		}
	}

	return e.Func == nil || e.Func.simpleExpr()
}

func (e CallExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.Func, singleLine, "call expression requires a function expression")
	w.WriteByte('(')
	e.Args.writeExprs(w, singleLine)
	w.WriteByte(')')
}

func (e IndexExpr) simpleExpr() bool {
	return (e.Slice == nil || e.Slice.simpleExpr()) &&
		(e.Index == nil || e.Index.simpleExpr())
}

func (e IndexExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.Slice, singleLine, "index expression requires a slice expression")
	w.WriteByte('[')
	writeExpr(w, e.Index, singleLine, "index expression requires an index expression")
	w.WriteByte(']')
}

func (e RangeExpr) simpleExpr() bool {
	return (e.Slice == nil || e.Slice.simpleExpr()) &&
		(e.Min == nil || e.Min.simpleExpr()) &&
		(e.Max == nil || e.Max.simpleExpr())
}

func (e RangeExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.Slice, singleLine, "index expression requires a slice expression")
	w.WriteByte('[')
	writeExpr(w, e.Min, singleLine, "")
	w.WriteByte(':')
	writeExpr(w, e.Max, singleLine, "")
	w.WriteByte(']')
}

func (e IdentExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e IdentExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('+')
	writeExpr(w, e.Op, singleLine, "identity expression requires an operand expression")
}

func (e NegateExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e NegateExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('-')
	writeExpr(w, e.Op, singleLine, "negation expression requires an operand expression")
}

func (e NotExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e NotExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('!')
	writeExpr(w, e.Op, singleLine, "not expression requires an operand expression")
}

func (e ComplementExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e ComplementExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('^')
	writeExpr(w, e.Op, singleLine, "complement expression requires an operand expression")
}

func (e AddrOfExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e AddrOfExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('&')
	writeExpr(w, e.Op, singleLine, "address-of expression requires an operand expression")
}

func (e DerefExpr) simpleExpr() bool {
	return (e.Op == nil || e.Op.simpleExpr())
}

func (e DerefExpr) writeExpr(w *code.Writer, singleLine bool) {
	w.WriteByte('*')
	writeExpr(w, e.Op, singleLine, "dereference expression requires an operand expression")
}

func (e AddExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e AddExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "add expression requires a left-hand operand expression")
	w.WriteString(" + ")
	writeExpr(w, e.RHS, singleLine, "add expression requires a right-hand operand expression")
}

func (e SubtractExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e SubtractExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "subtract expression requires a left-hand operand expression")
	w.WriteString(" - ")
	writeExpr(w, e.RHS, singleLine, "subtract expression requires a right-hand operand expression")
}

func (e MultiplyExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e MultiplyExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "multiply expression requires a left-hand operand expression")
	w.WriteString(" * ")
	writeExpr(w, e.RHS, singleLine, "multiply expression requires a right-hand operand expression")
}

func (e DivideExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e DivideExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "divide expression requires a left-hand operand expression")
	w.WriteString(" / ")
	writeExpr(w, e.RHS, singleLine, "divide expression requires a right-hand operand expression")
}

func (e ModulusExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e ModulusExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "modulus expression requires a left-hand operand expression")
	w.WriteString(" % ")
	writeExpr(w, e.RHS, singleLine, "modulus expression requires a right-hand operand expression")
}

func (e ShiftLeftExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e ShiftLeftExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "shift-left expression requires a left-hand operand expression")
	w.WriteString(" << ")
	writeExpr(w, e.RHS, singleLine, "shift-left expression requires a right-hand operand expression")
}

func (e ShiftRightExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e ShiftRightExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "shift-right expression requires a left-hand operand expression")
	w.WriteString(" >> ")
	writeExpr(w, e.RHS, singleLine, "shift-right expression requires a right-hand operand expression")
}

func (e EqualExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e EqualExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "equal expression requires a left-hand operand expression")
	w.WriteString(" == ")
	writeExpr(w, e.RHS, singleLine, "equal expression requires a right-hand operand expression")
}

func (e NotEqualExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e NotEqualExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "not-equal expression requires a left-hand operand expression")
	w.WriteString(" != ")
	writeExpr(w, e.RHS, singleLine, "not-equal expression requires a right-hand operand expression")
}

func (e LessThanExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e LessThanExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "less-than expression requires a left-hand operand expression")
	w.WriteString(" < ")
	writeExpr(w, e.RHS, singleLine, "less-than expression requires a right-hand operand expression")
}

func (e LessOrEqualExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e LessOrEqualExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "less-or-equal expression requires a left-hand operand expression")
	w.WriteString(" <= ")
	writeExpr(w, e.RHS, singleLine, "less-or-equal expression requires a right-hand operand expression")
}

func (e MoreThanExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e MoreThanExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "less-than expression requires a left-hand operand expression")
	w.WriteString(" > ")
	writeExpr(w, e.RHS, singleLine, "less-than expression requires a right-hand operand expression")
}

func (e MoreOrEqualExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e MoreOrEqualExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "less-or-equal expression requires a left-hand operand expression")
	w.WriteString(" >= ")
	writeExpr(w, e.RHS, singleLine, "less-or-equal expression requires a right-hand operand expression")
}

func (e BitAndExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e BitAndExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "bit-and expression requires a left-hand operand expression")
	w.WriteString(" & ")
	writeExpr(w, e.RHS, singleLine, "bit-and expression requires a right-hand operand expression")
}

func (e BitXorExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e BitXorExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "bit-xor expression requires a left-hand operand expression")
	w.WriteString(" ^ ")
	writeExpr(w, e.RHS, singleLine, "bit-xor expression requires a right-hand operand expression")
}

func (e BitOrExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e BitOrExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "bit-or expression requires a left-hand operand expression")
	w.WriteString(" | ")
	writeExpr(w, e.RHS, singleLine, "bit-or expression requires a right-hand operand expression")
}

func (e LogAndExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e LogAndExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "log-and expression requires a left-hand operand expression")
	w.WriteString(" && ")
	writeExpr(w, e.RHS, singleLine, "log-and expression requires a right-hand operand expression")
}

func (e LogOrExpr) simpleExpr() bool {
	return (e.LHS == nil || e.LHS.simpleExpr()) &&
		(e.RHS == nil || e.RHS.simpleExpr())
}

func (e LogOrExpr) writeExpr(w *code.Writer, singleLine bool) {
	writeExpr(w, e.LHS, singleLine, "log-or expression requires a left-hand operand expression")
	w.WriteString(" || ")
	writeExpr(w, e.RHS, singleLine, "log-or expression requires a right-hand operand expression")
}

func writeExpr(w *code.Writer, e Expr, singleLine bool, reqMessage string) {
	if e == nil {
		if reqMessage != "" {
			panic(errors.New(reqMessage))
		}

		return
	}

	e.writeExpr(w, singleLine)
}

func exprString(e Expr, reqMessage string) string {
	return writeString(func(w *code.Writer) { writeExpr(w, e, true, reqMessage) })
}

func (e Exprs) simpleExprs() bool {
	for _, itm := range e {
		if itm != nil && !itm.simpleExpr() {
			return false
		}
	}

	return true
}

func (e Exprs) writeExprs(w *code.Writer, singleLine bool) {
	if !singleLine {
		// Make an attempt to fit a single line
		singleLine = true

		for _, itm := range e {
			if itm != nil && !itm.simpleExpr() {
				singleLine = false
			}
		}
	}

	if singleLine {
		for idx, itm := range e {
			if idx > 0 {
				w.WriteString(", ")
			}

			writeExpr(w, itm, singleLine, "expression in list must not be nil")
		}
	} else {
		w.Newline()
		w.Indent(func(w *code.Writer) {
			for _, itm := range e {
				writeExpr(w, itm, singleLine, "expression in list must not be nil")
				w.WriteByte(',')
				w.Newline()
			}
		})
	}
}

var (
	_ Expr = Symbol{}
	_ Expr = NilExpr{}
	_ Expr = IotaExpr{}
	_ Expr = BoolExpr(false)
	_ Expr = IntExpr(0)
	_ Expr = UintExpr(0)
	_ Expr = FloatExpr(0)
	_ Expr = RuneExpr('r')
	_ Expr = StringExpr("")
	_ Expr = ParExpr{}
	_ Expr = CastExpr{}
	_ Expr = SliceExpr{}
	_ Expr = MapExpr{}
	_ Expr = StructExpr{}
	_ Expr = FuncExpr{}
	_ Expr = NewExpr{}
	_ Expr = MakeExpr{}
	_ Expr = MemberExpr{}
	_ Expr = CallExpr{}
	_ Expr = IndexExpr{}
	_ Expr = RangeExpr{}
	_ Expr = IdentExpr{}
	_ Expr = NegateExpr{}
	_ Expr = NotExpr{}
	_ Expr = ComplementExpr{}
	_ Expr = AddrOfExpr{}
	_ Expr = DerefExpr{}
	_ Expr = AddExpr{}
	_ Expr = SubtractExpr{}
	_ Expr = MultiplyExpr{}
	_ Expr = DivideExpr{}
	_ Expr = ModulusExpr{}
	_ Expr = ShiftLeftExpr{}
	_ Expr = ShiftRightExpr{}
	_ Expr = EqualExpr{}
	_ Expr = NotEqualExpr{}
	_ Expr = LessThanExpr{}
	_ Expr = LessOrEqualExpr{}
	_ Expr = MoreThanExpr{}
	_ Expr = MoreOrEqualExpr{}
	_ Expr = BitAndExpr{}
	_ Expr = BitXorExpr{}
	_ Expr = BitOrExpr{}
	_ Expr = LogAndExpr{}
	_ Expr = LogOrExpr{}
)
