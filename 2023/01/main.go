package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string

	digits = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	part1Re = digitFinder{
		first: regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(digits[:9], "|"))),
		last:  regexp.MustCompile(fmt.Sprintf(`.*(%s)`, strings.Join(digits[:9], "|"))),
	}
	part2Re = digitFinder{
		first: regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(digits, "|"))),
		last:  regexp.MustCompile(fmt.Sprintf(`.*(%s)`, strings.Join(digits, "|"))),
	}
)

type digitFinder struct {
	first, last *regexp.Regexp
}

func findDigit(line string, re *regexp.Regexp) int {
	matches := re.FindStringSubmatch(line)
	if len(matches) > 1 {
		return slices.Index(digits, matches[1])%9 + 1
	}
	return 0
}

func parse(line string, df digitFinder) int {
	first := findDigit(line, df.first)
	last := findDigit(line, df.last)
	return first*10 + last
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		calibration := parse(line, part1Re)
		res += calibration
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		calibration := parse(line, part2Re)
		res += calibration
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
