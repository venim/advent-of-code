package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string

	digits      = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitsRegex = regexp.MustCompile(fmt.Sprintf(`^(%s)`, strings.Join(digits, "|")))
)

func matchDigit(line string) int {
	if matches := digitsRegex.FindStringSubmatch(line); len(matches) > 1 {
		return slices.Index(digits, matches[1]) + 1
	}
	return 0
}

func parse(line string, withDigits bool) int {
	first := 0
	last := 0
	for i, s := range line {
		v, _ := strconv.Atoi(string(s))
		if withDigits {
			v = matchDigit(line[i:])
		}
		if v != 0 {
			if first == 0 {
				first = v
			}
			last = v
		}
	}
	return first*10 + last
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		calibration := parse(line, false)
		res += calibration
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		calibration := parse(line, true)
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
