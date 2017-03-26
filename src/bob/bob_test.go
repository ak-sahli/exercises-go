package bob

import "testing"

func TestHey(t *testing.T) {
	for _, testCase := range bobTestCases {
		observed := Hey(testCase.phrase)
		if observed != testCase.expected {
			t.Fatalf("observed: %s, expected: %s", observed, testCase.expected)
		}
	}
}
