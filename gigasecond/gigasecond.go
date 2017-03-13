package gigasecond

import "time"

func AddGigasecond(date time.Time) time.Time {
	return date.Add(1e9 * time.Second)
}
