package code_test

import (
	"fmt"
	"testing"

	code "github.com/trwk76/gocode"
)

func TestWriter(t *testing.T) {
	for _, item := range writerTests {
		item.test(t)
	}
}

type (
	writerTest struct {
		f   code.WriteFunc
		res string
	}
)

func (i writerTest) test(t *testing.T) {
	if res, err := code.WriteString("", i.f); err != nil {
		t.Error(err)
	} else if res != i.res {
		t.Errorf("expected:\n%s\ngot:\n%s\n", i.res, res)
	}
}

var writerTests []writerTest = []writerTest{
	{
		f: func(w *code.Writer) {
			fmt.Fprintf(w, `type %s struct {`, "MyStruct")
			w.Newline()

			w.Indent(func(w *code.Writer) {
				w.Table(
					code.TableRow{Columns: []string{"Object"}},
					code.TableRow{Columns: []string{"ID", "ID", "`json:\"id\"`"}},
					code.TableRow{Columns: []string{"Name", "string", "`json:\"name\"`"}},
				)
			})

			w.WriteByte('}')
			w.Newline()
		},
		res: "type MyStruct struct {\n\tObject\n\tID     ID     `json:\"id\"`\n\tName   string `json:\"name\"`\n}\n",
	},
}
