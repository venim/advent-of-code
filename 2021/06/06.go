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

func shift(state []int) []int {
	state[7] += state[0]
	return append(state[1:], state[0])
}
func solve(line string, days int) (res int) {
	state := make([]int, 9)
	for _, n := range strings.Split(line, ",") {
		state[mustAtoi(n)]++
	}
	for i := 0; i < days; i++ {
		state = shift(state)
	}
	for _, n := range state {
		res += n
	}
	return
}

func part1(lines []string) (res int) {
	return solve(lines[0], 80)
}

func part2(lines []string) (res int) {
	return solve(lines[0], 256)
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
