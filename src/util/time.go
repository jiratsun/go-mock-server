package timeutil

import "time"

func OfSeconds(value int64) time.Duration {
	return time.Second * time.Duration(value)
}

func OfMillis(value int64) time.Duration {
	return time.Millisecond * time.Duration(value)
}
