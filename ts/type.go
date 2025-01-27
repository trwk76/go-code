package ts

import "github.com/trwk76/go-code"

var (
	Any     AnyType
	Boolean BooleanType
	Number  NumberType
	String  StringType
)

type (
	Type interface {
		writeType(w *code.Writer)
	}

	AnyType     struct{}
	BooleanType struct{}
	NumberType  struct{}
	StringType  struct{}

	ArrayType struct {
		Item Type
	}
)

func (AnyType) writeType(w *code.Writer)     { w.WriteString(string(KwAny)) }
func (BooleanType) writeType(w *code.Writer) { w.WriteString(string(KwBoolean)) }
func (NumberType) writeType(w *code.Writer)  { w.WriteString(string(KwNumber)) }
func (StringType) writeType(w *code.Writer)  { w.WriteString(string(KwString)) }

func (t ArrayType) writeType(w *code.Writer) {
	t.Item.writeType(w)
	w.WriteString("[]")
}

var (
	_ Type = AnyType{}
	_ Type = BooleanType{}
	_ Type = NumberType{}
	_ Type = StringType{}
	_ Type = ArrayType{}
)
