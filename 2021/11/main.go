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

type octopus struct {
	level     int
	neighbors map[pos]*octopus
}

type board struct {
	SizeX, SizeY int
	Flashed      map[pos]struct{}
	Octopuses    map[pos]*octopus
}

func (b *board) addEnergy(octopuses map[pos]*octopus) {
	for p := range octopuses {
		b.Octopuses[p].level++
		if b.Octopuses[p].level > 9 {
			b.flash(p)
		}
	}
}

func (b *board) flash(p pos) {
	if _, ok := b.Flashed[p]; ok {
		return
	}
	b.Flashed[p] = struct{}{}
	b.addEnergy(b.Octopuses[p].neighbors)
}

func (b *board) step() {
	b.Flashed = map[pos]struct{}{}
	b.addEnergy(b.Octopuses)
	for p := range b.Flashed {
		b.Octopuses[p].level = 0
	}
}

func (b *board) neighbors(p pos) map[pos]*octopus {
	neighbors := map[pos]*octopus{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				p := pos{p.X + i, p.Y + j}
				if p.X >= 0 && p.X < b.SizeX &&
					p.Y >= 0 && p.Y < b.SizeY {
					neighbors[p] = nil
				}
			}
		}
	}
	return neighbors
}

func parse(lines []string) *board {
	b := &board{
		Octopuses: map[pos]*octopus{},
	}
	b.SizeY = len(lines)
	b.SizeX = len(lines[0])
	for y := 0; y < b.SizeY; y++ {
		for x := 0; x < b.SizeX; x++ {
			p := pos{x, y}
			neighbors := b.neighbors(p)
			level := mustAtoi(string(lines[y][x]))
			b.Octopuses[p] = &octopus{
				level:     level,
				neighbors: neighbors,
			}
		}
	}
	return b
}

func part1(lines []string) (res int) {
	b := parse(lines)
	for step := 0; step < 100; step++ {
		b.step()
		res += len(b.Flashed)
	}
	return
}

func part2(lines []string) (res int) {
	b := parse(lines)
	step := 0
	for {
		step++
		b.step()
		if len(b.Flashed) == b.SizeX*b.SizeY {
			return step
		}
	}
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
