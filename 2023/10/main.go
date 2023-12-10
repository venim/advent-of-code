package main

import (
	_ "embed"
	"flag"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

var pipes = map[byte][]util.Pos{
	'-': {util.Pos{X: -1, Y: 0}, util.Pos{X: 1, Y: 0}},
	'|': {util.Pos{X: 0, Y: -1}, util.Pos{X: 0, Y: 1}},
	'7': {util.Pos{X: -1, Y: 0}, util.Pos{X: 0, Y: 1}},
	'F': {util.Pos{X: 1, Y: 0}, util.Pos{X: 0, Y: 1}},
	'J': {util.Pos{X: -1, Y: 0}, util.Pos{X: 0, Y: -1}},
	'L': {util.Pos{X: 1, Y: 0}, util.Pos{X: 0, Y: -1}},
}

type node struct {
	distance int
	pos      util.Pos
}

type grid struct {
	lines []string
	start util.Pos
	parts map[util.Pos]int
}

func (g *grid) findStart() {
	for r := 0; r < len(g.lines); r++ {
		for c := 0; c < len(g.lines[0]); c++ {
			if g.lines[r][c] == 'S' {
				g.start = util.Pos{X: c, Y: r}
				return
			}
		}
	}
}

func (g *grid) seed() []node {
	queue := []node{}
	for _, d := range util.Directions {
		pos := util.Pos{X: g.start.X + d.X, Y: g.start.Y + d.Y}
		if pos.Y < 0 || pos.Y >= len(g.lines) || pos.X < 0 || pos.X >= len(g.lines[0]) {
			continue
		}
		for _, p := range pipes[g.lines[pos.Y][pos.X]] {
			next := util.Pos{X: pos.X + p.X, Y: pos.Y + p.Y}
			if next == g.start {
				queue = append(queue, node{1, pos})
			}
		}
	}
	g.replaceS(queue)

	return queue
}

func (g *grid) replaceS(seed []node) {
loop:
	for k, v := range pipes {
		for _, p := range seed {
			pos := util.Pos{X: p.pos.X - g.start.X, Y: p.pos.Y - g.start.Y}
			if !slices.Contains(v, pos) {
				continue loop
			}
		}
		line := []byte(g.lines[g.start.Y])
		line[g.start.X] = k
		g.lines[g.start.Y] = string(line)
		return
	}
}

func (g *grid) traverse() {
	visited := map[util.Pos]int{g.start: -1}
	queue := g.seed()

	for len(queue) > 0 {
		cur := queue[0]
		pos := cur.pos

		queue = queue[1:]

		visited[pos] = cur.distance
		for _, d := range pipes[g.lines[pos.Y][pos.X]] {
			pos := util.Pos{X: pos.X + d.X, Y: pos.Y + d.Y}
			if pos.Y < 0 || pos.Y >= len(g.lines) || pos.X < 0 || pos.X >= len(g.lines[0]) {
				continue
			}
			if visited[pos] == 0 {
				queue = append(queue, node{cur.distance + 1, pos})
			}
		}
	}
	g.parts = visited
}

func part1(lines []string) (res int) {
	g := &grid{lines: lines}
	g.findStart()
	g.traverse()

	for _, d := range g.parts {
		res = max(res, d)
	}

	return
}

func (g *grid) countIn() (res int) {
	for r := 0; r < len(g.lines); r++ {
		hParity := false
		for c := 0; c < len(g.lines[0]); c++ {
			pos := util.Pos{X: c, Y: r}
			if g.parts[pos] != 0 {
				if strings.Contains("|LJ", string(g.lines[r][c])) {
					hParity = !hParity
				}
				continue
			}

			if hParity {
				res++
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	g := &grid{lines: lines}
	g.findStart()
	g.traverse()

	return g.countIn()
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
