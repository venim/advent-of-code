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
		score := .5
		for _, n := range got {
			if want[n] {
				score *= 2
			}
		}
		res += int(score)
	}
	return
}

func part2(lines []string) (res int) {
	scratchcards := map[string]int{}
	for i, line := range lines {
		scratchcards[line]++
		want, got := parseCard(line)
		j := i + 1
		for _, n := range got {
			if want[n] {
				if j >= len(lines) {
					continue
				}
				scratchcards[lines[j]] += scratchcards[line]
				j++
			}
		}
	}
	for _, n := range scratchcards {
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
