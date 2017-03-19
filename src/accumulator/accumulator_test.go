package accumulator

import "testing"

func TestAccumulate(t *testing.T) {
	for _, testCase := range accumulateTestCases {
		expected := testCase.expected
		observed := Accumulate(testCase.given, testCase.converter)
		for index, expectedWord := range expected {
			observedWord := observed[index]
			if expectedWord != observedWord {
				t.Fatalf("observed: %s, expected: %s", observedWord, expectedWord)
			}
		}
	}
}
