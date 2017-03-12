package clock

import "fmt"

const (
	format         = "%02d:%02d"
	minutesPerHour = 60
	hoursPerDay    = 24
	minutesPerDay  = 60 * hoursPerDay
)

type Clock struct {
	hours, minutes int
}

func NewClock(hours, minutes int) Clock {
	normalizedHours, normalizedMinutes := normalize(hours, minutes)
	return Clock{normalizedHours, normalizedMinutes}
}

func (clock Clock) AddMinutes(minutesToAdd int) Clock {
	return NewClock(clock.hours, clock.minutes+minutesToAdd)
}

func (clock Clock) String() string {
	return fmt.Sprintf(format, clock.hours, clock.minutes)
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
