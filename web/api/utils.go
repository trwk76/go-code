package api

import "fmt"

func uniqueKey[T any](m map[string]T, base string, def string) string {
	if base == "" {
		base = def
	}

	_, fnd := m[base]
	if !fnd {
		return base
	}

	i := 1
	name := fmt.Sprintf("%s%d", base, i)
	for _, fnd = m[name]; fnd; {
		i++
		name = fmt.Sprintf("%s%d", base, i)
	}

	return name
}