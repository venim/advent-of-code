package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
	lines []string

	oppGuide = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	part1Guide = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	part2Guide = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
)

func scoreGame(opp, you int) int {
	score := you
	if opp == you {
		score += 3
	} else {
		won := opp < you
		// flip result if rock vs. scissors
		if opp+you == 4 {
			won = !won
		}
		if won {
			score += 6
		}
	}
	return score
}

func guessGame(opp int, result int) int {
	for i := 1; i <= 3; i++ {
		if score := scoreGame(opp, i); score == result+i {
			return score
		}
	}
	return 0
}

func parsePart2(line string) (int, int) {
	moves := strings.Split(line, " ")
	return oppGuide[moves[0]], part2Guide[moves[1]]
}

func parsePart1(line string) (int, int) {
	moves := strings.Split(line, " ")
	return oppGuide[moves[0]], part1Guide[moves[1]]
}

func part2() {
	sum := 0
	for _, line := range lines {
		sum += guessGame(parsePart2(line))
	}
	fmt.Printf("[Part 2] = %d\n", sum)
}

func part1() {
	sum := 0
	for _, line := range lines {
		sum += scoreGame(parsePart1(line))
	}
	fmt.Printf("[Part 1] = %d\n", sum)
}

func main() {
	lines = strings.Split(input, "\n")

	var t time.Time

	t = time.Now()
	part1()
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	part2()
	fmt.Printf("took %s\n\n", time.Since(t))
}
