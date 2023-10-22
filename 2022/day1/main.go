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

func init() {
	// do this in init (not main) so test file has same input
	// input = strings.TrimRight(input, "\n")
	// if len(input) == 0 {
	// panic("empty input.txt file")
	// }
}

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
