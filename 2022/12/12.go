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
	input      string
	directions = []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

type pos [2]int

type node struct {
	pos
	dist int
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func findPath(grid [][]byte, starts []pos, end pos) int {
	seen := make(map[pos]struct{})
	q := make([]node, 0, len(starts))
	for _, n := range starts {
		q = append(q, node{n, 0})
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.pos == end {
			return n.dist
		}
		if _, ok := seen[n.pos]; ok {
			continue
		}
		seen[n.pos] = struct{}{}
		var x, y, dx, dy int
		x = n.pos[0]
		y = n.pos[1]
		for _, dir := range directions {
			dx = dir[0]
			dy = dir[1]
			if 0 <= x+dx && x+dx < len(grid) &&
				0 <= y+dy && y+dy < len(grid[0]) &&
				int(grid[x+dx][y+dy])-int(grid[x][y]) <= 1 {
				q = append(q, node{pos{x + dx, y + dy}, n.dist + 1})
			}
		}
	}
	return -1
}

func makeGrid(lines []string) (grid [][]byte, start pos, end pos, as []pos) {
	grid = make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]byte, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			grid[i][j] = lines[i][j]
			switch lines[i][j] {
			case byte('S'):
				grid[i][j] = 'a'
				start = pos{i, j}
			case byte('E'):
				grid[i][j] = 'z'
				end = pos{i, j}
			case byte('a'):
				as = append(as, pos{i, j})
			}
		}
	}
	return
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()
	var (
		t time.Time
	)
	lines := strings.Split(input, "\n")
	grid, start, end, as := makeGrid(lines)

	t = time.Now()
	part1 := findPath(grid, []pos{start}, end)
	glog.Infof("[Part 1] = %v", part1)
	glog.Infof("took %s\n\n", time.Since(t))

	t = time.Now()
	part2 := findPath(grid, as, end)
	glog.Infof("[Part 2] = %v", part2)
	glog.Infof("took %s\n\n", time.Since(t))
}
