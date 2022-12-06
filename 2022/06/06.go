package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

var (
	//go:embed in.txt
	input string
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getMarker(stream string, n int) int {

	for i := 0; i < len(stream)-n; i++ {
		found := map[byte]bool{}
		for j := i; j < i+n; j++ {
			if _, ok := found[stream[j]]; !ok {
				found[stream[j]] = true
			} else {
				break
			}
		}
		if len(found) == n {
			return i + n
		}
	}
	return -1
}

func part1(stream string) (res int) {
	return getMarker(stream, 4)
}

func part2(stream string) (res int) {
	return getMarker(stream, 14)
}

func main() {
	var (
		t   time.Time
		res int
	)

	t = time.Now()
	res = part1(input)
	fmt.Printf("[Part 1] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	res = part2(input)
	fmt.Printf("[Part 2] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))
}
