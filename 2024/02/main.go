package main

import (
	_ "embed"
	"flag"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func toIntSlice(s []string) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = util.MustAtoi(s[i])
	}
	return res
}

func isSafe(levels []int) bool {
	if !slices.IsSorted(levels) {
		slices.Reverse(levels)
		if !slices.IsSorted(levels) {
			slices.Reverse(levels)
			return false
		}
		slices.Reverse(levels)
	}
	for i := 1; i < len(levels); i++ {
		delta := levels[i] - levels[i-1]
		d := math.Abs(float64(delta))
		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}

func part1(lines []string) (res int) {
	for _, l := range lines {
		if isSafe(toIntSlice(strings.Fields(l))) {
			res++
		}
	}
	return
}

func removeLevel(l []int, i int) []int {
	return slices.Delete(slices.Clone(l), i, i+1)
}

func part2(lines []string) (res int) {
line:
	for _, l := range lines {
		levels := toIntSlice(strings.Fields(l))
		if isSafe(levels) {
			res++
			continue
		}
		for i := 0; i < len(levels); i++ {
			l := removeLevel(levels, i)
			if isSafe(l) {
				res++
				continue line
			}
		}
	}
	return
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
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
