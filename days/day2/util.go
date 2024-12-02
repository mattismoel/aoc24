package day2

import (
	"math"
)

const (
	MIN_INCREASE = 1
	MAX_INCREASE = 3
)

func isAscending(a int, b int) bool {
	return a-b <= 0
}

func isDescending(a int, b int) bool {
	return a-b >= 0
}

func isSafeDiff(a int, b int) bool {
	diff := math.Abs(float64(a) - float64(b))
	if diff < MIN_INCREASE || diff > MAX_INCREASE {
		return false
	}

	return true
}
