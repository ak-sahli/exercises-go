package leap

import "testing"

func TestLeapYears(t *testing.T) {
	for _, testCase := range testCases {
		observed := Leap(testCase.year)
		if observed != testCase.expected {
			t.Fatalf("%d, observed: %t, expected: %t", testCase.year, observed, testCase.expected)
		}
	}
}