package acronym

var abbreviateTestCases = []struct {
	phrase, expected string
} {
	{"Portable Network Graphics", "PNG"},
	{"HyperText Markup Language", "HTML"},
	{"Ruby on Rails", "ROR"},
	{"PHP: Hypertext Preprocessor", "PHP"},
	{"First In, First Out", "FIFO"},
	{"Complementary metal-oxide semiconductor", "CMOS"},
}