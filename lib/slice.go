package lib

import "fmt"

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
