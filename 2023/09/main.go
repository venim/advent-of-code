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

type history struct {
	Sequences [][]int
}

func diff(s []int) []int {
	o := make([]int, 0, len(s)-1)
	for i := 1; i < len(s); i++ {
		o = append(o, s[i]-s[i-1])
	}
	return o
}

func allZero(s []int) bool {
	for _, n := range s {
		if n != 0 {
			return false
		}
	}
	return true
}

func generate(line string, reverse bool) history {
	h := history{}
	l := numRe.FindAllString(line, -1)
	s := make([]int, 0, len(l))
	for _, n := range l {
		s = append(s, util.MustAtoi(n))
	}
	if reverse {
		slices.Reverse(s)
	}
	h.Sequences = append(h.Sequences, s)
	for {
		s = diff(s)
		if len(s) == 0 {
			break
		}
		h.Sequences = append(h.Sequences, s)

		if allZero(s) {
			break
		}
	}
	return h
}

func (h history) extrapolate() {
	h.Sequences[len(h.Sequences)-1] = append(h.Sequences[len(h.Sequences)-1], 0)

	for i := len(h.Sequences) - 2; i >= 0; i-- {
		p := h.Sequences[i+1]
		s := h.Sequences[i]
		h.Sequences[i] = append(s, s[len(s)-1]+p[len(p)-1])
	}
}

func part1(lines []string) (res int) {
	for _, l := range lines {
		h := generate(l, false)
		h.extrapolate()
		s := h.Sequences[0]
		res += s[len(s)-1]
	}
	return
}

func part2(lines []string) (res int) {
	for _, l := range lines {
		h := generate(l, true)
		h.extrapolate()
		s := h.Sequences[0]
		res += s[len(s)-1]
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
