package main

import (
	_ "embed"
	"flag"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
	numRe = regexp.MustCompile(`(\d+)`)
)

type Maps []Map

func (ms Maps) convert(s int) int {
	for _, m := range ms {
		s = m.convert(s)
	}
	return s
}

type Map struct {
	rules []Rule
}

func (m Map) convert(s int) int {
	for _, mr := range m.rules {
		if mr.src <= s && s < mr.src+mr.length {
			return s - mr.src + mr.dst
		}
	}
	return s
}

type Rule struct {
	dst    int
	src    int
	length int
}

func parse(lines []string) (string, Maps) {
	seeds := lines[0]

	// read rules
	ms := []Map{}
	lines = lines[1:]

	for len(lines) > 0 {
		m := Map{}
		lines = lines[1:]
		for len(lines) > 0 && lines[0] != "" {
			parts := numRe.FindAllString(lines[0], -1)
			lines = lines[1:]

			if len(parts) == 0 {
				continue
			}
			r := Rule{
				dst:    util.MustAtoi(parts[0]),
				src:    util.MustAtoi(parts[1]),
				length: util.MustAtoi(parts[2]),
			}
			m.rules = append(m.rules, r)
		}
		ms = append(ms, m)
	}
	return seeds, ms
}

func part1(lines []string) (res int) {
	seeds, ms := parse(lines)

	// convert each seed
	minLocation := math.MaxInt
	for _, seed := range numRe.FindAllString(seeds, -1) {
		seed := util.MustAtoi(seed)
		minLocation = min(minLocation, ms.convert(seed))
	}

	return minLocation
}

func part2(lines []string) (res int) {
	seeds, ms := parse(lines)

	seedRanges := numRe.FindAllString(seeds, -1)
	minLocation := math.MaxInt

	for i := 0; i < len(seedRanges); i += 2 {
		start := util.MustAtoi(seedRanges[i])
		length := util.MustAtoi(seedRanges[i+1])
		for seed := start; seed < start+length; seed++ {
			minLocation = min(minLocation, ms.convert(seed))
		}
	}
	return minLocation
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t   time.Time
		res any
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
