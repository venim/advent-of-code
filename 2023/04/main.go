package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

var numRe = regexp.MustCompile(`(\d+)`)

func parseCard(line string) (map[string]bool, []string) {
	line = strings.Split(line, ": ")[1]
	parts := strings.Split(line, "|")
	want := map[string]bool{}
	for _, n := range numRe.FindAllString(parts[0], -1) {
		want[n] = true
	}
	got := numRe.FindAllString(parts[1], -1)
	return want, got
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		want, got := parseCard(line)
		// glog.Infof("%v | %v", want, got)
		score := 0
		for _, got := range got {
			if want[got] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		res += score
	}
	return
}

func part2(lines []string) (res int) {
	lineCopies := map[string]int{}
	for i, line := range lines {
		lineCopies[line]++
		want, got := parseCard(line)
		matched := 0
		for _, got := range got {
			if want[got] {
				matched++
			}
		}
		for j := i + 1; j <= i+matched; j++ {
			if j > len(lines) {
				break
			}
			lineCopies[lines[j]] += lineCopies[line]
		}
	}
	for _, n := range lineCopies {
		res += n
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
