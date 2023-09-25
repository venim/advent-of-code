package main

import (
	_ "embed"
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

func part1(lines []string) (res int) {
	prev := -1
	for _, _depth := range lines {
		depth := mustAtoi(_depth)
		if depth > prev && prev != -1 {
			res++
		}
		prev = depth
	}
	return
}

func part2(lines []string) (res int) {
	sum := 0
	sum += mustAtoi(lines[0])
	sum += mustAtoi(lines[1])
	sum += mustAtoi(lines[2])
	prev := sum
	for i := 3; i < len(lines); i++ {
		sum -= mustAtoi(lines[i-3])
		sum += mustAtoi(lines[i])
		if sum > prev {
			res++
		}
		prev = sum
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
