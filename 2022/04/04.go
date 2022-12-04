package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed in.txt
var input string

func parseSections(sections string) (int64, int64) {
	parts := strings.Split(sections, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)
	return start, end
}

func containsSet(a, b string) bool {
	i, j := parseSections(a)
	x, y := parseSections(b)
	return i <= x && j >= y || x <= i && y >= j
}

func setsOverlap(a, b string) bool {
	i, j := parseSections(a)
	x, y := parseSections(b)
	return i <= y && x <= j
}

func part2() {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sets := strings.Split(line, ",")
		if setsOverlap(sets[0], sets[1]) {
			sum++
		}
	}
	fmt.Printf("[Part 2] = %d\n", sum)
}

func part1() {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sets := strings.Split(line, ",")
		if containsSet(sets[0], sets[1]) {
			sum++
		}
	}
	fmt.Printf("[Part 1] = %d\n", sum)
}

func main() {
	var t time.Time

	t = time.Now()
	part1()
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	part2()
	fmt.Printf("took %s\n\n", time.Since(t))
}
