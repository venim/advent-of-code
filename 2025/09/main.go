package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func area(a, b util.Pos) int {
	x := abs(a.X-b.X) + 1
	y := abs(a.Y-b.Y) + 1

	return x * y
}

func part1(lines []string) (res int) {
	var tiles []util.Pos
	for _, l := range lines {
		p := strings.Split(l, ",")
		tiles = append(tiles, util.Pos{X: util.MustAtoi(p[0]), Y: util.MustAtoi(p[1])})
	}

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			area := area(tiles[i], tiles[j])
			res = max(area, res)
		}
	}

	return
}

func area2(a, b util.Pos, valid map[util.Pos]bool) int {
	if valid[util.Pos{a.X, b.Y}] && valid[util.Pos{b.X, a.Y}] {
		x := abs(a.X-b.X) + 1
		y := abs(a.Y-b.Y) + 1

		return x * y
	}
	return 0
}

func printValid(x, y int, valid map[util.Pos]bool) {
	for j := 0; j <= y+1; j++ {
		for i := 0; i <= x+1; i++ {
			if valid[util.Pos{i, j}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func part2(lines []string) (res int) {
	maxX, maxY := 0, 0
	var tiles []util.Pos
	for _, l := range lines {
		parts := strings.Split(l, ",")
		pos := util.Pos{X: util.MustAtoi(parts[0]), Y: util.MustAtoi(parts[1])}
		maxX = max(maxX, pos.X)
		maxY = max(maxY, pos.Y)
		tiles = append(tiles, pos)
	}

	valid := map[util.Pos]bool{}
	rows := make([]*util.Pos, maxY+2)

	for i := 0; i < len(tiles); i++ {
		a := tiles[i]
		b := tiles[0]
		if i+1 != len(tiles) {
			b = tiles[i+1]
		}

		if a.X == b.X {
			for i := min(a.Y, b.Y); i <= max(a.Y, b.Y); i++ {
				valid[util.Pos{a.X, i}] = true
			}
		} else {
			for i := min(a.X, b.X); i <= max(a.X, b.X); i++ {
				valid[util.Pos{i, a.Y}] = true
			}
		}
	}

	for t := range valid {
		if rows[t.Y] == nil {
			rows[t.Y] = &util.Pos{t.X, t.X}
		} else {
			rows[t.Y].X = min(rows[t.Y].X, t.X)
			rows[t.Y].Y = max(rows[t.Y].Y, t.X)
		}
	}

	for y, r := range rows {
		if r == nil {
			continue
		}
		for x := r.X; x <= r.Y; x++ {
			valid[util.Pos{x, y}] = true
		}
	}

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			area := area2(tiles[i], tiles[j], valid)
			res = max(area, res)
		}
	}

	return
}

func init() {
	_ = flag.Set("logtostderr", "true")
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
