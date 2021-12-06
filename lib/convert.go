package lib

import "strconv"

// ToInt Basically ignoring errors
func ToInt(value string) (result int) {
	result, _ = strconv.Atoi(value)
	return
}

// ToIntBin Basically ignoring errors and converting from int64 to int
func ToIntBin(value string) int {
	result64, _ := strconv.ParseInt(value, 2, 64)
	return int(result64)
}
