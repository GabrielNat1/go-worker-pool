package utils

import "time"

// SleepMs sleeps milliseconds
func SleepMs(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
