package utils

import (
	"math/rand/v2"
	"time"
)

func GetRandomDailyDelay() time.Duration {
	// Random delay between 0 and 24 hours
	hours := rand.IntN(3) + 24
	minutes := rand.IntN(60)
	seconds := rand.IntN(60)

	return time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second
}