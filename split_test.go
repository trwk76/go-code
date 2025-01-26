package code_test

import (
	"slices"
	"strings"
	"testing"

	code "github.com/trwk76/go-code"
)

func TestSplit(t *testing.T) {
	for _, item := range splitTests {
		item.test(t)
	}
}

type (
	splitTest struct {
		id  string
		res []string
	}
)

func (i splitTest) test(t *testing.T) {
	res := code.SplitID(i.id)

	if !slices.Equal(res, i.res) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", strings.Join(i.res, ", "), strings.Join(res, ", "))
	}
}

var splitTests []splitTest = []splitTest{
	{
		id:  "__XMLName__",
		res: []string{"XML", "Name"},
	},
	{
		id:  "_xmlName__",
		res: []string{"xml", "Name"},
	},
	{
		id:  "XML_Name",
		res: []string{"XML", "Name"},
	},
	{
		id:  "  XML_Name  ",
		res: []string{"XML", "Name"},
	},
}
