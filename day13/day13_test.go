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

	expected := 17
	result := Part1(input, 15, 11)
	if result != expected {
		t.Fatalf(`Part1() = %d, want %d, error`, result, expected)
	}
}
