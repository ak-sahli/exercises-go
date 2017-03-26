package bob

import (
	"strings"
	"unicode"
)

func Hey(greetings string) string {
	greetings = strings.TrimSpace(greetings)
	if greetings == "" {
		return "Fine. Be that way!"
	}
	if any(greetings, unicode.IsUpper) && !any(greetings, unicode.IsLower) {
		return "Whoa, chill out!"
	}
	if greetings[len(greetings)-1] == '?' {
		return "Sure."
	}
	return "Whatever."
}

func any(phrase string, predicate func(rune) bool) bool {
	for _, symbol := range phrase {
		if predicate(symbol) {
			return true
		}
	}
	return false
}
