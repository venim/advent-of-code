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

func makeSet(sections string) map[int64]bool {
	set := make(map[int64]bool, 10)
	parts := strings.Split(sections, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)
	for i := start; i <= end; i++ {
		set[i] = true
	}
	return set
}
func compareSize(a, b map[int64]bool) (big, small map[int64]bool) {
	if len(a) > len(b) {
		big = a
		small = b
	} else {
		small = a
		big = b
	}
	return big, small
}

func containsSet(a, b map[int64]bool) bool {
	big, small := compareSize(a, b)

	for i := range small {
		if _, ok := big[i]; !ok {
			return false
		}
	}
	return true
}

func setsOverlap(a, b map[int64]bool) bool {
	big, small := compareSize(a, b)

	for i := range small {
		if _, ok := big[i]; ok {
			return true
		}
	}
	return false
}

func part2() {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sets := strings.Split(line, ",")
		if setsOverlap(makeSet(sets[0]), makeSet(sets[1])) {
			sum++
		}
	}
	fmt.Printf("[Part 2] = %d\n", sum)
}

func part1() {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sets := strings.Split(line, ",")
		if containsSet(makeSet(sets[0]), makeSet(sets[1])) {
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
