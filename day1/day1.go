package main

import (
	"fmt"
	"github.com/charlesbourget/aoc2021/lib"
	"math"
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
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	return count
}

func part2(input []int) int {
	count := 0
	lastSum := math.MaxInt
	for i := 0; i < len(input)-2; i++ {
		if input[i+2] > lastSum {
			count++
		}
		lastSum = input[i]
	}

	return count
}
