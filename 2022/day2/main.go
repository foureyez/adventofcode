package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	win  = 6
	loss = 0
	draw = 3
)

var (
	oppChoices  = map[rune]rune{'A': 'R', 'B': 'P', 'C': 'S'}
	myChoices   = map[rune]rune{'X': 'R', 'Y': 'P', 'Z': 'S'}
	choiceScore = map[rune]int{'R': 1, 'P': 2, 'S': 3}
)

func init() {
	input = strings.TrimRight(input, "\n")
}

// Link: https://adventofcode.com/2022/day/2
func main() {
	ans := part1(input)
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	out := strings.Split(input, "\n")
	score := 0
	for _, line := range out {
		choices := []rune(line)
		myChoice := myChoices[choices[2]]
		oppChoice := oppChoices[choices[0]]
		score += choiceScore[myChoice]

		if oppChoice == myChoice {
			score += draw
		} else if (myChoice == 'R' && oppChoice == 'P') ||
			(myChoice == 'P' && oppChoice == 'S') ||
			(myChoice == 'S' && oppChoice == 'R') {
			// score += loss
		} else {
			score += win
		}
	}
	return score
}
