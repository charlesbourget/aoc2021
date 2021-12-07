package lib

import (
	"fmt"
	"math"
)

// Sum To sum all values of int slice
func Sum(values []int) (count int) {
	for _, v := range values {
		count += v
	}

	return
}

// Unpack To unpack a slice of two elements
func Unpack(values []string) (string, string) {
	if len(values) != 2 {
		fmt.Println("Unpack only accept len=2 this will fail...")
		return "", ""
	}

	return values[0], values[1]
}

// UnpackInt To unpack a slice of two elements and convert them to int
func UnpackInt(values []string) (int, int) {
	value1, value2 := Unpack(values)

	return ToInt(value1), ToInt(value2)
}

// ToIntSlice Convert a slice of string to a slice of int
func ToIntSlice(values []string) []int {
	result := make([]int, len(values))
	for i, v := range values {
		result[i] = ToInt(v)
	}

	return result
}

func Min(values []int) (min int) {
	min, _ = MinMax(values)
	return
}

func Max(values []int) (max int) {
	_, max = MinMax(values)
	return
}

func MinMax(values []int) (min int, max int) {
	min = math.MaxInt
	for _, v := range values {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return
}

func Seq(min int, max int) []int {
	r := make([]int, 0, max-min+1)
	for v := min; v <= max; v++ {
		r = append(r, v)
	}

	return r
}
