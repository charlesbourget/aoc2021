package main

import (
	"fmt"
	"github.com/charlesbourget/aoc2021/lib"
	"strconv"
	"strings"
)

func main() {
	input, err := lib.Read("day6/input")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}

func Part1(input []string) int {
	days := parse(input[0])

	for day := 0; day < 80; day++ {
		days = simulate(days)
	}

	return countFish(days)
}

func Part2(input []string) int {
	days := parse(input[0])

	for day := 0; day < 256; day++ {
		days = simulate(days)
	}

	return countFish(days)
}

func parse(input string) []int {
	numbers := strings.Split(input, ",")
	days := make([]int, 9)
	for _, number := range numbers {
		daysToHatch, _ := strconv.Atoi(number)
		days[daysToHatch] += 1
	}
	return days
}

func simulate(days []int) []int {
	newSlice := make([]int, 9)
	for i, v := range days {
		if i == 0 {
			newSlice[8] = v
			newSlice[6] = v
		} else {
			newSlice[i-1] += v
		}
	}

	return newSlice
}

func countFish(days []int) (count int) {
	for _, v := range days {
		count += v
	}

	return
}
