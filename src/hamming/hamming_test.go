package hamming

import "testing"

func TestDistance(t *testing.T) {
	for _, testCase := range distanceTestCases {
		observedDistance, error := Distance(testCase.strand1, testCase.strand2)
		expectedDistance := testCase.expected
		if observedDistance < 0 && error == nil {
			t.Fatal("Error cannot be nil when strands have different sizes")
		}
		if observedDistance != expectedDistance {
			t.Fatalf("observed: %d, expected: %d", observedDistance, expectedDistance)
		}
	}
}
