package accumulator

func Accumulate(words []string, convertFunction func(string) string) []string {
	for index, word := range words {
		words[index] = convertFunction(word)
	}
	return words
}
