package main

import (
	"fmt"
	"github.com/charlesbourget/aoc2021/lib"
	"strconv"
)

func main() {
	input, err := lib.Read("day3/input")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}

func Part1(input []string) int64 {
	acc := make([]int, len(input[0]))
	for _, v := range input {
		chars := []rune(v)
		for i := 0; i < len(chars); i++ {
			acc[i] += int(chars[i] - '0')
		}
	}
	gamma := ""
	eps := ""
	for _, v := range acc {
		if v < (len(input) / 2) {
			eps += "1"
			gamma += "0"
		} else {
			eps += "0"
			gamma += "1"
		}
	}
	gammaValue, _ := strconv.ParseInt(gamma, 2, 64)
	epsValue, _ := strconv.ParseInt(eps, 2, 64)
	return gammaValue * epsValue
}

func Part2(input []string) int64 {
	oxyRating, _ := strconv.ParseInt(findRating(input, 0, true), 2, 64)
	co2Rating, _ := strconv.ParseInt(findRating(input, 0, false), 2, 64)
	return oxyRating * co2Rating
}

func findRating(data []string, position int, isOxy bool) string {
	if len(data) == 1 {
		return data[0]
	}

	acc := 0
	value := 0
	for _, v := range data {
		chars := []rune(v)
		acc += int(chars[position] - '0')
		if acc > (len(data)-1)/2 {
			value = 1
			break
		}
	}

	if !isOxy {
		if value == 1 {
			value = 0
		} else {
			value = 1
		}
	}

	var filteredData []string
	for _, v := range data {
		chars := []rune(v)
		if int(chars[position]-'0') == value {
			filteredData = append(filteredData, v)
		}
	}

	return findRating(filteredData, position+1, isOxy)
}
