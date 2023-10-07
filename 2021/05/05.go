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

func parseLine(line string) (start pos, end pos) {
	points := strings.Split(line, " -> ")
	p1 := strings.Split(points[0], ",")
	p2 := strings.Split(points[1], ",")
	start = pos{mustAtoi(p1[0]), mustAtoi(p1[1])}
	end = pos{mustAtoi(p2[0]), mustAtoi(p2[1])}
	return
}

func solve(lines []string, shouldDraw func(pos, pos) bool) (res int) {
	board := map[pos]int{}
	for _, line := range lines {
		start, end := parseLine(line)
		if !shouldDraw(start, end) {
			continue
		}
		d := &pos{}

		if start.Y < end.Y {
			d.Y = 1
		} else if start.Y > end.Y {
			d.Y = -1
		}
		if start.X < end.X {
			d.X = 1
		} else if start.X > end.X {
			d.X = -1
		}

		for {
			if _, ok := board[start]; !ok {
				board[start] = 0
			}
			board[start]++
			start.X += d.X
			start.Y += d.Y
			if d.X != 0 && start.X-d.X == end.X ||
				d.Y != 0 && start.Y-d.Y == end.Y {
				break
			}
		}
	}

	for _, n := range board {
		if n > 1 {
			res++
		}
	}
	return
}

func part1(lines []string) (res int) {
	return solve(lines, func(start, end pos) bool {
		return start.X == end.X || start.Y == end.Y
	})
}

func part2(lines []string) (res int) {
	return solve(lines, func(_, _ pos) bool {
		return true
	})
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
