package leap

var testCases = []struct {
	year        int
	expected    bool
}{
	{2015, false},
	{2016, true},
	{2100, false},
	{2000, true},
}
