package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

type point struct {
	X int
	Y int
}

func (p *point) diff(p2 *point) (dx int, dy int) {
	dx = clamp(p.X - p2.X)
	dy = clamp(p.Y - p2.Y)
	return
}

func (p *point) drop(g *grid) bool {
	p.Y++
	if !g.blocked(p) {
		return true
	}
	p.X--
	if !g.blocked(p) {
		return true
	}
	p.X += 2
	if !g.blocked(p) {
		return true
	}
	p.Y--
	p.X--
	return false
}

func clamp(i int) int {
	if i < 0 {
		return -1
	}
	if i > 0 {
		return 1
	}
	return 0
}

type grid struct {
	Rocks map[point]struct{}
	Sand  map[point]struct{}
	Draw  bool
	Depth int
	MaxX  int
	MinX  int
}

func newGrid() *grid {
	g := &grid{
		Rocks: make(map[point]struct{}),
		Sand:  make(map[point]struct{}),
		Draw:  false,
		Depth: 0,
		MaxX:  500,
		MinX:  500,
	}
	return g
}

func (g *grid) makeRocks(line string) {
	points := strings.Split(line, " -> ")
	var start *point
	for _, s := range points {
		parts := strings.Split(s, ",")
		p := &point{mustAtoi(parts[0]), mustAtoi(parts[1])}
		if p.Y > g.Depth {
			g.Depth = p.Y
		}
		if p.X < g.MinX {
			g.MinX = p.X
		}
		if p.X > g.MaxX {
			g.MaxX = p.X
		}
		if start == nil {
			start = p
		}
		g.Rocks[*p] = struct{}{}
		dx, dy := p.diff(start)
		for !(start.X == p.X && start.Y == p.Y) {
			start.X += dx
			start.Y += dy
			g.Rocks[*start] = struct{}{}
		}
	}
	if g.Draw {
		g.MaxX += 20
		g.MinX -= 20
	}
}

func (g *grid) dropSand(s *point, depth int) bool {
	for s.drop(g) {
		if s.Y == depth {
			return false
		}
	}
	if g.Draw {
		g.draw()
	}
	g.Sand[*s] = struct{}{}
	if s.X == 500 && s.Y == 0 {
		return false
	}
	return true
}

func (g *grid) blocked(s *point) (ok bool) {
	if s.Y == g.Depth+2 {
		return true
	}
	if _, ok = g.Rocks[*s]; ok {
		return true
	}
	_, ok = g.Sand[*s]
	return ok
}

func (g *grid) draw() {
	fmt.Printf("Grains: %d\n", len(g.Sand))
	var line []byte
	for y := 0; y <= g.Depth+2; y++ {
		line = make([]byte, 0, g.MaxX-g.MinX+1)
		for x := g.MinX; x <= g.MaxX; x++ {
			p := point{x, y}
			if _, ok := g.Rocks[p]; ok {
				line = append(line, byte('#'))
			} else if _, ok := g.Sand[p]; ok {
				line = append(line, byte('o'))
			} else {
				line = append(line, byte('.'))
			}
		}
		line = append(line, byte('\n'))
		os.Stdout.Write(line)
	}
	time.Sleep(25 * time.Millisecond)
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part1(lines []string) (res int) {
	g := newGrid()
	for _, line := range lines {
		g.makeRocks(line)
	}
	for g.dropSand(&point{500, 0}, g.Depth) {
		res++
	}
	return
}

func part2(lines []string) (res int) {
	g := newGrid()
	for _, line := range lines {
		g.makeRocks(line)
	}
	if g.Draw {
		for x := g.MinX; x <= g.MaxX; x++ {
			g.Rocks[point{x, g.Depth + 2}] = struct{}{}
		}
	}
	for g.dropSand(&point{500, 0}, g.Depth+2) {
		res++
	}
	return res + 1
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
