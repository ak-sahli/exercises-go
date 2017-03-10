package leap

import "testing"

var testCases = []struct {
	year        int
	expected    bool
}{
	{2015, false},
	{2016, true},
	{2100, false},
	{2000, true},
}

func TestLeapYears(t *testing.T) {
	for _, testCase := range testCases {
		observed := Leap(testCase.year)
		if observed != testCase.expected {
			t.Fatalf("%d, observed: %t, expected: %t", testCase.year, observed, testCase.expected)
		}
	}
}