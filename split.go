package code

import (
	"strings"
	"unicode"
)

// SplitID splits an identifier into words that seem relevant based on the casing or _ separators.
func SplitID(id string) []string {
	res := make([]string, 0)
	runes := []rune(id)

	for len(runes) > 0 {
		if part, ok := readIDPart(&runes); ok {
			res = append(res, part)
		}
	}

	return res
}

func readIDPart(r *[]rune) (string, bool) {
	// Remove leading non-identifier characters.
	for len(*r) > 0 && !isIDChar((*r)[0]) {
		*r = (*r)[1:]
	}

	for len(*r) > 0 && isIDSep((*r)[0]) {
		*r = (*r)[1:]
	}

	if len(*r) < 1 {
		return "", false
	}

	buf := strings.Builder{}

	buf.WriteRune((*r)[0])
	*r = (*r)[1:]
	prevLower := false

	for len(*r) > 0 && !isIDPartEnd(*r, prevLower) {
		prevLower = false

		if isIDChar((*r)[0]) {
			// Make sure to leave out any trailing noise
			buf.WriteRune((*r)[0])
			prevLower = unicode.IsLower((*r)[0])
		}

		*r = (*r)[1:]
	}

	return buf.String(), true
}

func isIDPartEnd(r []rune, prevLower bool) bool {
	if len(r) < 1 {
		return true
	}

	if isIDSep(r[0]) {
		return true
	} else if prevLower && unicode.IsUpper(r[0]) {
		return true
	} else if len(r) > 1 && unicode.IsUpper(r[0]) && unicode.IsLower(r[1]) {
		return true
	}

	return false
}

func isIDChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || isIDSep(r)
}

func isIDSep(r rune) bool {
	return r == '_'
}
