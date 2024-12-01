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

func part1(lines []string) (res int) {
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)
		left = append(left, util.MustAtoi(parts[0]))
		right = append(right, util.MustAtoi(parts[1]))
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := 0; i < len(lines); i++ {
		res += int(math.Abs(float64(left[i] - right[i])))
	}
	return
}

func part2(lines []string) (res int) {
	left := make([]int, 0, len(lines))
	right := make(map[int]int)
	for _, line := range lines {
		parts := strings.Fields(line)
		left = append(left, util.MustAtoi(parts[0]))
		right[util.MustAtoi(parts[1])]++
	}
	for _, n := range left {
		if m, ok := right[n]; ok {
			res += n * m
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
