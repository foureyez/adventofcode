package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

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
	out := strings.Split(input, "\n")
	finalScore := 0
	for _, items := range out {
		midPoint := len(items) / 2
		firstCompartmentMap := make(map[rune]any)
		for i := 0; i < midPoint; i++ {
			item := rune(items[i])
			firstCompartmentMap[item] = nil
		}

		for i := midPoint; i < len(items); i++ {
			item := rune(items[i])
			_, ok := firstCompartmentMap[item]
			if ok {
				if unicode.IsLower(item) {
					finalScore += int(item) - 96
				} else {
					finalScore += int(item) - 38
				}
				break
			}
		}
	}
	return finalScore
}

func part2(input string) int {
	return 0
}
