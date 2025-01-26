package code_test

import (
	"testing"

	code "github.com/trwk76/go-code"
)

func TestCase(t *testing.T) {
	for _, item := range caseTests {
		item.test(t)
	}
}

type (
	caseTest struct {
		f   code.IDTransformer
		id  string
		res string
	}
)

func (i caseTest) test(t *testing.T) {
	res := i.f(i.id)

	if res != i.res {
		t.Errorf("expected:\n%s\ngot:\n%s\n", i.res, res)
	}
}

var caseTests []caseTest = []caseTest{
	{
		f:   code.IDToCamel,
		id:  "__XMLName__",
		res: "xmlName",
	},
	{
		f:   code.IDToPascal,
		id:  "__XMLName__",
		res: "XmlName",
	},
	{
		f:   code.IDToSnake,
		id:  "__XMLName__",
		res: "xml_name",
	},
}
