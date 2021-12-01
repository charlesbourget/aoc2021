package main

import (
	"fmt"
	"github.com/charlesbourget/aoc2021/lib"
)

func main() {
	input, err := lib.ReadInt("day1/input")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []int) int {
	count := 0
	for i := range input {
		if i > 0 && input[i-1] < input[i] {
			count++
		}
	}
	return count
}

func part2(input []int) int {
	count := 0
	lastSum := 0

	for i := range input {
		if i+2 < len(input) {
			sum := input[i+2] + input[i+1] + input[i]
			if sum > lastSum && lastSum > 0 {
				count++
			}
			lastSum = sum
		}
	}

	return count
}
