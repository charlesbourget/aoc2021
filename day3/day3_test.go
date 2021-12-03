package main

import (
	"fmt"
	"github.com/charlesbourget/aoc2021/lib"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := lib.Read("input.test")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	var expected int64 = 198
	result := Part1(input)
	if result != expected {
		t.Fatalf(`Part1() = %d, want %d, error`, result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := lib.Read("input.test")
	if err != nil {
		fmt.Println("Error while reading input. ", err)
		return
	}

	var expected int64 = 230
	result := Part2(input)
	if result != expected {
		t.Fatalf(`Part2() = %d, want %d, error`, result, expected)
	}
}
