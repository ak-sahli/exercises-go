package acronym

import "testing"

func TestAbbreviate(t *testing.T) {
	for _, testCase := range abbreviateTestCases {
		observed := Abbreviate(testCase.phrase)
		expected := testCase.expected
		if observed != expected {
			t.Fatalf("observed: %s, expected: %s", observed, expected)
		}
	}
}

