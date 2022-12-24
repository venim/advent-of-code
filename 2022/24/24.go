package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func (p pos) valid(gs pos, bs map[pos]struct{}) bool {
	// not valid if the position is a blizzard
	if _, ok := bs[p]; ok {
		return false
	}
	// pos in grid
	return p.X >= 0 &&
		p.X < gs.X &&
		p.Y >= 0 &&
		p.X < gs.Y
}

type blizzard struct {
	Pos    pos
	Dir    pos
	Symbol byte
}

func (b *blizzard) move(max pos) {
	b.Pos.X += b.Dir.X
	b.Pos.Y += b.Dir.Y

	if b.Pos.X < 0 {
		b.Pos.X = max.X - 1
	} else {
		b.Pos.X %= max.X
	}
	if b.Pos.Y < 0 {
		b.Pos.Y = max.Y - 1
	} else {
		b.Pos.Y %= max.Y
	}
}

type grid struct {
	Blizzards []*blizzard
	Size      pos
}

func (g grid) traverse(start, end pos) (steps int) {
	states := make(map[pos]struct{})
	states[start] = struct{}{}
	// draw(g.Blizzards, g.Size, states)
	for {
		// new minute
		steps++
		// glog.Info(res, len(states))

		// move blizzards and record all new locations
		blizzardLocations := map[pos]struct{}{}
		for _, b := range g.Blizzards {
			b.move(g.Size)
			blizzardLocations[b.Pos] = struct{}{}
		}

		nextStates := make(map[pos]struct{})
		// for each current potential state
		for state := range states {
			// for each potential next location
			for _, d := range dirs {
				next := pos{state.X + d.X, state.Y + d.Y}
				// did we find the end?
				if next == end {
					return
				}
				// add to next step if inside the grid and not a blizzard
				if next.valid(g.Size, blizzardLocations) {
					nextStates[next] = struct{}{}
				}
			}
			// can we wait in our spot?
			if _, ok := blizzardLocations[state]; !ok {
				nextStates[state] = struct{}{}
			}
		}
		states = nextStates
		// draw(g.Blizzards, g.Size, states)
	}
}

func parse(lines []string) (g grid, start pos, end pos) {
	g.Size = pos{len(lines) - 2, len(lines[0]) - 2}

	for j := 1; j < len(lines[0])-1; j++ {
		if lines[0][j] == '.' {
			start = pos{-1, j - 1}
		}
		if lines[len(lines)-1][j] == '.' {
			end = pos{len(lines) - 2, j - 1}
		}
	}

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == '.' {
				continue
			}
			g.Blizzards = append(g.Blizzards, &blizzard{Pos: pos{i - 1, j - 1}, Dir: dirs[lines[i][j]], Symbol: lines[i][j]})
		}
	}
	return
}

func draw(blizzards []*blizzard, gridSize pos, states map[pos]struct{}) {
	grid := [][]byte{}
	var row []byte
	for i := 0; i < gridSize.X; i++ {
		row = []byte{}
		for j := 0; j < gridSize.Y; j++ {
			var (
				p = pos{i, j}
				n int
				s byte
			)
			for _, b := range blizzards {
				if b.Pos == p {
					n++
					s = b.Symbol
				}
			}
			if n > 1 {
				row = append(row, byte(n+48))
			} else if n == 0 {

				for s := range states {
					if p == s {
						n++
						row = append(row, byte('E'))
						break
					}
				}
				if n == 0 {
					row = append(row, byte('.'))
				}
			} else {
				row = append(row, s)
			}
		}
		grid = append(grid, row)
	}
	var s string
	for _, r := range grid {
		s += fmt.Sprintf("%s\n", r)
	}
	glog.Info("\n" + s)
}

func contains(states map[pos]struct{}, end pos) bool {
	for s := range states {
		if s == end {
			return true
		}
	}
	return false
}

func part1(lines []string) (res int) {
	g, start, end := parse(lines)
	return g.traverse(start, end)
}

func part2(lines []string) (res int) {
	g, start, end := parse(lines)
	res += g.traverse(start, end)
	res += g.traverse(end, start)
	res += g.traverse(start, end)
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
