package pangram

import "unicode"

func IsPangram(phrase string) bool {
	alphabet := make(map[rune]bool, 26)
	for letter := 'a'; letter <= 'z'; letter++ {
		alphabet[letter] = false
	}
	for _, symbol := range []rune(phrase) {
		if unicode.IsLetter(symbol) {
			letter := unicode.ToLower(symbol)
			alphabet[letter] = true
		}
	}

	for _, letter := range alphabet {
		if !letter {
			return false
		}
	}
	return true
}
