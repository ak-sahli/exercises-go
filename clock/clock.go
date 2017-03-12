package clock

import "fmt"

const (
	FORMAT         = "%02d:%02d"
	MINUTES_P_HOUR = 60
	HOURS_P_DAY    = 24
	MINUTES_P_DAY  = 60 * HOURS_P_DAY
)

type Clock struct {
	hours, minutes int
}

func NewClock(hours, minutes int) Clock {
	normalizedHours, normalizedMinutes := normalize(hours, minutes)
	return Clock{normalizedHours, normalizedMinutes}
}

func (clock Clock) String() string {
	return fmt.Sprintf(FORMAT, clock.hours, clock.minutes)
}

func toMinutes(hours, minutes int) int {
	return hours*MINUTES_P_HOUR + minutes
}

func normalize(hours, minutes int) (int, int) {
	totalMinutes := toMinutes(hours, minutes)
	if totalMinutes < 0 {
		totalMinutes = (totalMinutes % MINUTES_P_DAY) + MINUTES_P_DAY
	}
	return totalMinutes / MINUTES_P_HOUR % HOURS_P_DAY, totalMinutes % MINUTES_P_HOUR
}
