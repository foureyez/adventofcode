package main

import (
	_ "embed"
	"flag"
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
	lossChoice  = map[rune]rune{'R': 'S', 'P': 'R', 'S': 'P'}
	winChoice   = map[rune]rune{'R': 'P', 'P': 'S', 'S': 'R'}
)

func init() {
	input = strings.TrimRight(input, "\n")
}

// Link: https://adventofcode.com/2022/day/2
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
	score := 0
	for _, line := range out {
		currGame := []rune(line)
		myChoice := myChoices[currGame[2]]
		oppChoice := oppChoices[currGame[0]]
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

func part2(input string) int {
	out := strings.Split(input, "\n")
	score := 0
	for _, line := range out {
		currGame := []rune(line)
		oppChoice := oppChoices[currGame[0]]
		if currGame[2] == 'X' {
			myChoice := lossChoice[oppChoice]
			score += choiceScore[myChoice]
		} else if currGame[2] == 'Y' {
			score += draw + choiceScore[oppChoice]
		} else {
			myChoice := winChoice[oppChoice]
			score += win + choiceScore[myChoice]
		}
	}
	return score
}
