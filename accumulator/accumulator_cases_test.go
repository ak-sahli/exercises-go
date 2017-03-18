package accumulator

import "strings"

var accumulateTestCases = []struct {
	given, expected []string
	converter       func(string) string
}{
	{
		[]string{},
		[]string{},
		echo,
	},
	{
		[]string{"echo", "echo", "echo", "echo"},
		[]string{"echo", "echo", "echo", "echo"},
		echo,
	},
	{
		[]string{"first", "letter", "only"},
		[]string{"First", "Letter", "Only"},
		capitalize,
	},
	{
		[]string{"hello", "world"},
		[]string{"HELLO", "WORLD"},
		upper,
	},
}

func echo(word string) string {
	return word
}

func capitalize(word string) string {
	return strings.Title(word)
}

func upper(word string) string {
	return strings.ToUpper(word)
}
