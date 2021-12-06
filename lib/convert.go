package lib

import "strconv"

// ToInt Basically ignoring errors
func ToInt(value string) (result int) {
	result, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return
}

// ToIntBin Basically ignoring errors and converting from int64 to int
func ToIntBin(value string) int {
	result64, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(result64)
}
