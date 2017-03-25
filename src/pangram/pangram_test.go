package pangram

import "testing"

func TestIsPangram(t *testing.T) {
	for _, testCase := range isPangramTestCases {
		observed := IsPangram(testCase.phrase)
		expected := testCase.expected
		if observed != expected {
			t.Fatalf("observed: %t, expected: %t", observed, expected)
		}
	}
}
