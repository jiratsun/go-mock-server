package time

import "time"

func OfSeconds(value int) time.Duration {
	return time.Second * time.Duration(value)
}

func OfMillis(value int) time.Duration {
	return time.Millisecond * time.Duration(value)
}
