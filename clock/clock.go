package clock

import "fmt"

const (
	format         = "%02d:%02d"
	minutesPerHour = 60
	hoursPerDay    = 24
	minutesPerDay  = 60 * hoursPerDay
)

type Time struct {
	hours, minutes int
}

func NewClock(hours, minutes int) Time {
	normalizedHours, normalizedMinutes := normalize(hours, minutes)
	return Time{normalizedHours, normalizedMinutes}
}

func (time Time) AddMinutes(minutesToAdd int) Time {
	return NewClock(time.hours, time.minutes+minutesToAdd)
}

func (time Time) String() string {
	return fmt.Sprintf(format, time.hours, time.minutes)
}

func toMinutes(hours, minutes int) int {
	return hours*minutesPerHour + minutes
}

func normalize(hours, minutes int) (int, int) {
	totalMinutes := toMinutes(hours, minutes)
	if totalMinutes < 0 {
		totalMinutes = (totalMinutes % minutesPerDay) + minutesPerDay
	}
	return totalMinutes / minutesPerHour % hoursPerDay, totalMinutes % minutesPerHour
}
