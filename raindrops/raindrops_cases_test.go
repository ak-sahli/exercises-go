package raindrops

var soudTestCases = []struct {
	number   int
	expected string
}{
	{1, "1"},
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
	{6, "Pling"},
	{9, "Pling"},
	{10, "Plang"},
	{14, "Plong"},
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
}
