package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
)

func parseInventory(inventory string) []int64 {
	elves := []int64{}
	lines := strings.Split(inventory, "\n\n")
	for _, line := range lines {
		var sum int64
		for _, food := range strings.Split(line, "\n") {
			cal, _ := strconv.ParseInt(food, 10, 64)
			sum += cal
		}
		elves = append(elves, sum)
	}
	return elves
}

func part2() {
	elves := parseInventory(input)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	sum := elves[0] + elves[1] + elves[2]
	fmt.Printf("[Part 2] = %d\n", sum)
}

func part1() {
	var max int64
	for _, elf := range parseInventory(input) {
		if elf > max {
			max = elf
		}
	}
	fmt.Printf("[Part 1] = %d\n", max)
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
