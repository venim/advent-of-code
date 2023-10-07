package main

import (
	_ "embed"
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"golang.org/x/exp/slices"
)

var (
	//go:embed in.txt
	input string
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

func solve(lines []string, fuelFn func(int, int) int) int {
	// sort crab list
	line := strings.Split(lines[0], ",")
	crabs := make([]int, 0, len(line))
	for _, n := range line {
		crabs = append(crabs, mustAtoi(n))
	}
	slices.Sort(crabs)
	// from lowest pos to highest
	minFuel := -1
	for i := crabs[0]; i <= crabs[len(crabs)-1]; i++ {
		// find fuel cost
		totalFuel := 0
		for _, c := range crabs {
			totalFuel += fuelFn(c, i)
		}
		// keep minimum
		if minFuel == -1 {
			minFuel = totalFuel
		} else {
			minFuel = min(minFuel, totalFuel)
		}
	}
	return minFuel
}

func part1(lines []string) (res int) {
	return solve(lines, func(c, i int) int {
		fuel := c - i
		if fuel < 0 {
			fuel *= -1
		}
		return fuel
	})
}

func part2(lines []string) (res int) {
	return solve(lines, func(c, i int) int {
		steps := c - i
		if steps < 0 {
			steps *= -1
		}
		fuel := 0
		for i := 0; i <= steps; i++ {
			fuel += i
		}
		return fuel
	})
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t   time.Time
		res int
	)
	flag.Parse()
	lines := strings.Split(input, "\n")

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
