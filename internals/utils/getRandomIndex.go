package utils

import "math/rand/v2"

func GenerateRandomIndex(length int) int {
	if length <= 0 {
		return -1 // handle empty slices
	}
	return rand.IntN(length)
}