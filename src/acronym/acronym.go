package acronym

import (
	"regexp"
	"strings"
)

const regEx = "((?:^|[A-Z]+|)[a-z:]+)"

func Abbreviate(phrase string) string {
	var abbreviation string
	words := regexp.MustCompile(regEx).FindAllString(phrase, -1)
	for word := range words {
		firstLetter := string(words[word][0])
		abbreviation += strings.ToUpper(firstLetter)
	}
	return abbreviation
}
