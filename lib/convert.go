package lib

import "strconv"

func ToInt(value string) (result int) {
	result, _ = strconv.Atoi(value)
	return
}

func ToIntBin(value string) int {
	result64, _ := strconv.ParseInt(value, 2, 64)
	return int(result64)
}
