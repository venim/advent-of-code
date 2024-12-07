package main

import (
	_ "embed"
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func findGuard(lines []string) util.Pos {
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] == '^' {
				return util.Pos{X: r, Y: c}
			}
		}
	}
	return util.Pos{}
}

var dirs = []byte{'^', '>', 'v', '<'}

func part1(lines []string) (res int) {
	pos := findGuard(lines)
	dir := 0
	visited := map[util.Pos]bool{pos: true}
	for {
		new := pos.Add(util.Directions[dirs[dir]])
		if new.IsOutOfBounds(len(lines), len(lines[0])) {
			return len(visited)
		}
		if lines[new.X][new.Y] == '#' {
			dir++
			dir %= 4
		} else {
			pos = new
			visited[pos] = true
		}
	}
}

type node struct {
	pos util.Pos
	dir byte
}

func isLoop(lines []string, pos util.Pos) bool {
	dir := 0
	visited := map[node]bool{}
	for {
		new := pos.Add(util.Directions[dirs[dir]])
		if new.IsOutOfBounds(len(lines), len(lines[0])) {
			return false
		}
		if lines[new.X][new.Y] == '#' {
			dir++
			dir %= 4
		} else {
			pos = new
			n := node{pos, dirs[dir]}
			if visited[n] {
				return true
			}
			visited[n] = true
		}
	}
}

func part2(lines []string) (res int) {
	start := findGuard(lines)
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] == '.' {
				tmp := lines[r]
				lines[r] = fmt.Sprintf("%s#%s", lines[r][:c], lines[r][c+1:])
				if isLoop(lines, start) {
					res++
				}
				lines[r] = tmp
			}
		}
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
	lines := util.SplitLines(input)

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
