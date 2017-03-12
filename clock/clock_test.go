package clock

import "testing"

func TestCreateClock(t *testing.T) {
	for _, testCase := range createClockCases {
		observed := NewClock(testCase.hours, testCase.minutes).String()
		if observed != testCase.expected {
			t.Fatalf("observed: %s, expected: %s", observed, testCase.expected)
		}
	}
}
