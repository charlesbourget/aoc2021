package lib

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinInt(x int, y int) int {
	min := math.Min(float64(x), float64(y))
	return int(min)
}
