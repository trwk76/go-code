package code

import (
	"strings"
	"unicode"
)

// IDToCamel converts an ID to camel case (ex: fooBar).
func IDToCamel(id string) string {
	items := SplitID(id)

	for idx, item := range items {
		if idx == 0 {
			items[idx] = strings.ToLower(item)
		} else {
			items[idx] = IDToPascal(item)
		}
	}

	return strings.Join(items, "")
}

// IDToPascal converts an ID to Pascal case (ex: FooBar).
func IDToPascal(id string) string {
	items := SplitID(id)

	for idx, item := range items {
		items[idx] = string(unicode.ToUpper(rune(item[0]))) + strings.ToLower(item[1:])
	}

	return strings.Join(items, "")
}

// IDToSnake converts an ID to snake case (ex: foo_bar).
func IDToSnake(id string) string {
	items := SplitID(id)

	for idx, item := range items {
		items[idx] = strings.ToLower(item)
	}

	return strings.Join(items, "_")
}

type IDTransformer func(id string) string
