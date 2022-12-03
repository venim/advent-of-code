package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	//go:embed in.txt
	input string
	lines []string
)

func intersection[T comparable](a, b []T) []T {
	set := map[T]bool{}
	res := []T{}
	for _, i := range a {
		set[i] = false
	}
	for _, i := range b {
		if dup, ok := set[i]; ok {
			if !dup {
				res = append(res, i)
				set[i] = true
			}
		}
	}
	return res
}

func getCommonItems(in string) []byte {
	return intersection([]byte(in[:len(in)/2]), []byte(in[len(in)/2:]))
}

func itemPriority(b byte) int {
	i := int(b)
	if unicode.IsLower(rune(b)) {
		return i - 'a' + 1
	} else if unicode.IsUpper((rune(b))) {
		return i - 'A' + 27
	}
	return 0
}

func getBadges(elves ...string) []byte {
	badges := []byte(elves[0])
	for i := 1; i < len(elves); i++ {
		badges = intersection(badges, []byte(elves[i]))
	}
	return badges
}

func score(common []byte) int {
	sum := 0
	for _, b := range common {
		sum += itemPriority(b)
	}
	return sum
}

func part1() {
	sum := 0
	for _, sack := range lines {
		sum += score(getCommonItems((sack)))
	}
	fmt.Printf("[Part 1] %d\n", sum)
}

func part2() {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		sum += score(getBadges(lines[i], lines[i+1], lines[i+2]))
	}
	fmt.Printf("[Part 2] %d\n", sum)
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

	os.Exit(0)
}
