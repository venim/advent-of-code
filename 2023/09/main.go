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

type sequence []int

func (s sequence) diff() sequence {
	d := make(sequence, len(s)-1)
	for i := 1; i < len(s); i++ {
		d[i-1] = s[i] - s[i-1]
	}
	return d
}

func (s sequence) allZero() bool {
	for _, n := range s {
		if n != 0 {
			return false
		}
	}
	return true
}

func (s sequence) extrapolate() int {
	if s.allZero() {
		return 0
	}
	return s[len(s)-1] + s.diff().extrapolate()
}

func parse(line string) sequence {
	ns := numRe.FindAllString(line, -1)
	s := make(sequence, len(ns))
	for i, n := range ns {
		s[i] = util.MustAtoi(n)
	}
	return s
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		s := parse(line)
		res += s.extrapolate()
	}
	return
}

func part2(lines []string) (res int) {
	for _, l := range lines {
		s := parse(l)
		slices.Reverse(s)
		res += s.extrapolate()
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
