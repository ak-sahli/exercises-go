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

func TestAddMinutes(t *testing.T) {
	for _, testCase := range addMinutesCase {
		observed := NewClock(testCase.hours, testCase.minutes).AddMinutes(testCase.minutesToAdd)
		if observed.String() != testCase.expected {
			t.Fatalf("observed: %s, expected: %s", observed, testCase.expected)
		}
	}
}
