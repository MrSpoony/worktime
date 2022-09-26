package diff

import "time"

// GetTimeDiff returns the duration from start to end minus the break
func GetTimeDiff(start, end, breakdur time.Duration) time.Duration {
	now := time.Now()
	futuretime := now.Add(end).Add(-start).Add(-breakdur)
	diff := futuretime.Sub(now)
	return diff
}

// IntsToTimeDiff returns a duration create from a slice of ints
// the integer slice has to be of length two {hours, minutes} seconds are not supported yet
func IntsToTimeDiff(tm []int) time.Duration {
	now := time.Now()
	duration := now.Add(time.Hour*time.Duration(tm[0]) + time.Minute*time.Duration(tm[1])).Sub(now)
	return duration
}
