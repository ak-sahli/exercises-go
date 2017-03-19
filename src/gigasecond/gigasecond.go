package gigasecond

import "time"

func AddGigaSecond(date time.Time) time.Time {
	return date.Add(1e9 * time.Second)
}
