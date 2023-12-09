package main

import (
	_ "embed"
	"flag"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
	numRe = regexp.MustCompile(`(-?\d+)`)
)

func diff(s []int) []int {
	o := make([]int, 0, len(s)-1)
	for i := 1; i < len(s); i++ {
		o = append(o, s[i]-s[i-1])
	}
	return o
}

func allZero(n []int) bool {
	for _, n := range n {
		if n != 0 {
			return false
		}
	}
	return true
}

func extrapolate(n []int) int {
	if allZero(n) {
		return 0
	}
	return n[len(n)-1] + extrapolate(diff(n))
}

func parse(line string) []int {
	l := numRe.FindAllString(line, -1)
	s := make([]int, 0, len(l))
	for _, n := range l {
		s = append(s, util.MustAtoi(n))
	}
	return s
}

func part1(lines []string) (res int) {
	for _, l := range lines {
		s := parse(l)
		res += extrapolate(s)
	}
	return
}

func part2(lines []string) (res int) {
	for _, l := range lines {
		s := parse(l)
		slices.Reverse(s)
		res += extrapolate(s)
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
