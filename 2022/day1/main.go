package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// URL: https://adventofcode.com/2022/day/1
func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

/**
Simpler solution would've been to just create a array of sum of calories and
Part1 -> Find max
Part2 -> Sort the array and find top 3 and sum.
**/

// Edge case missed here when the input doesn't end with newline(\n),
// and the max calories is carried by last elf.
func part1(input string) int {
	maxCalories := 0
	currCalories := 0

	out := strings.Split(input, "\n")
	for _, k := range out {
		if k == "" {
			if currCalories > maxCalories {
				maxCalories = currCalories
			}
			currCalories = 0
			continue
		}
		cal, _ := strconv.Atoi(k)
		currCalories += cal
	}
	return maxCalories
}

func part2(input string) int {
	var currCalories, caloriesFirst, caloriesSecond, caloriesThird int

	out := strings.Split(input, "\n")
	for _, k := range out {
		if k == "" {
			if currCalories > caloriesFirst {
				caloriesThird = caloriesSecond
				caloriesSecond = caloriesFirst
				caloriesFirst = currCalories
			} else if currCalories > caloriesSecond {
				caloriesThird = caloriesSecond
				caloriesSecond = currCalories
			} else if currCalories > caloriesThird {
				caloriesThird = currCalories
			}
			currCalories = 0
			continue
		}
		cal, _ := strconv.Atoi(k)
		currCalories += cal
	}

	return caloriesFirst + caloriesSecond + caloriesThird
}
