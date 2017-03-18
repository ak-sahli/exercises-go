package raindrops

import "testing"

func TestConvert(t *testing.T) {
	for _, testCase := range soudTestCases {
		observed := Convert(testCase.number)
		expected := testCase.expected
		if observed != expected {
			t.Fatalf("observed: %d, expected: %d", observed, expected)
		}
	}
}
