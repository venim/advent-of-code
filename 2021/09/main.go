package main

import (
	_ "embed"
	"flag"
	"slices"
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

func getV(lines []string, p pos) int {
	return mustAtoi(string(lines[p.Y][p.X]))
}

func isLowPoint(lines []string, p pos) bool {
	v := getV(lines, p)
	for _, d := range []pos{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		x := p.X + d.X
		y := p.Y + d.Y
		if x >= 0 &&
			x < len(lines[0]) &&
			y >= 0 &&
			y < len(lines) {
			v2 := getV(lines, pos{x, y})
			if v2 <= v {
				return false
			}
		}
	}

	return true
}

func part1(lines []string) (res int) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			p := pos{x, y}
			v := getV(lines, p)
			if isLowPoint(lines, p) {
				res += v + 1
			}
		}
	}
	return
}

type node struct {
	p pos
	v int
}

func findBasin(lines []string, p pos) (size int) {
	visited := map[pos]struct{}{}
	queue := []node{{p, -1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		p := cur.p
		v := getV(lines, p)

		// skip if we've already been here
		if _, ok := visited[p]; ok {
			continue
		}
		// skip if the basin doesn't continue
		if v <= cur.v || v == 9 {
			continue
		}
		// save our place and add to the basin size
		visited[p] = struct{}{}
		size++

		// add each unvisited neighbor to the queue
		for _, d := range []pos{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			x := p.X + d.X
			y := p.Y + d.Y
			if x >= 0 &&
				x < len(lines[0]) &&
				y >= 0 &&
				y < len(lines) {
				cur := node{pos{x, y}, v}
				if _, ok := visited[cur.p]; !ok {
					queue = append(queue, cur)
				}
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	basinSizes := []int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			p := pos{x, y}
			if isLowPoint(lines, p) {
				size := findBasin(lines, p)
				basinSizes = append(basinSizes, size)
			}
		}
	}
	slices.Sort(basinSizes)
	res = 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		res *= size
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
