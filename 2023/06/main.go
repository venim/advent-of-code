package main

import (
	_ "embed"
	"flag"
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

func solve(times []string, distances []string) (res int) {
	res = 1
	for raceN := 0; raceN < len(times); raceN++ {
		wins := 0
		time := util.MustAtoi(times[raceN])
		distance := util.MustAtoi(distances[raceN])
		speed := 0
		for i := 0; i < time; i++ {
			d := speed * (time - i)
			if d > distance {
				wins++
			}
			speed++
		}
		res *= wins
	}
	return
}

func part1(lines []string) (res int) {
	times := numRe.FindAllString(lines[0], -1)
	distances := numRe.FindAllString(lines[1], -1)

	return solve(times, distances)
}

func part2(lines []string) (res int) {
	times := numRe.FindAllString(strings.ReplaceAll(lines[0], " ", ""), -1)
	distances := numRe.FindAllString(strings.ReplaceAll(lines[1], " ", ""), -1)

	return solve(times, distances)
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
