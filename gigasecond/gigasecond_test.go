package gigasecond

import (
	"testing"
	"time"
)

const dateFormat = "2006-01-02T15:04:05"

func TestAddGigasecond(t *testing.T) {
	for _, testCase := range addGigascecondCases {
		date, expectedDate := parse(testCase.date, testCase.expectedDate, t)
		observed := AddGigasecond(date)
		if observed != expectedDate {
			t.Fatalf("date: %s, expectedDate: %s, observed: %s", date, expectedDate, observed)
		}
	}
}

func parse(date string, expected string, t *testing.T) (time.Time, time.Time) {
	parsedDate, parseDateError := time.Parse(dateFormat, date)
	parsedExpectedDate, parseExpectedDateError := time.Parse(dateFormat, expected)
	if parseDateError != nil {
		t.Fatalf("Invalid data in test case : %s", date)
	}
	if parseExpectedDateError != nil {
		t.Fatalf("Invalid data in test case : %s", expected)
	}
	return parsedDate, parsedExpectedDate
}
