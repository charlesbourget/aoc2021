package lib

import "fmt"

func Sum(values []int) (count int) {
	for _, v := range values {
		count += v
	}

	return
}

func Unpack(values []string) (string, string) {
	if len(values) != 2 {
		fmt.Println("Unpack only accept len=2 this will fail...")
	}

	return values[0], values[1]
}

func UnpackInt(values []string) (int, int) {
	value1, value2 := Unpack(values)

	return ToInt(value1), ToInt(value2)
}
