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

func part1(lines []string) (res int) {
	galaxies := []util.Pos{}

	emptyRows := map[int]struct{}{}
	emptyCols := map[int]struct{}{}
	seenRows := map[int]struct{}{}
	seenCols := map[int]struct{}{}
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] != '.' {
				seenRows[r] = struct{}{}
				seenCols[c] = struct{}{}
				galaxies = append(galaxies, util.Pos{X: c, Y: r})
			}
		}
	}
	for r := 0; r < len(lines); r++ {
		if _, ok := seenRows[r]; !ok {
			emptyRows[r] = struct{}{}
		}
	}
	for c := 0; c < len(lines[0]); c++ {
		if _, ok := seenCols[c]; !ok {
			emptyCols[c] = struct{}{}
		}
	}
	n := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			xs := []int{g1.X, g2.X}
			slices.Sort(xs)
			x := xs[1] - xs[0]
			for c := range emptyCols {
				if c > xs[0] && c < xs[1] {
					x++
				}
			}
			ys := []int{g1.Y, g2.Y}
			slices.Sort(ys)
			y := ys[1] - ys[0]
			for r := range emptyRows {
				if r > ys[0] && r < ys[1] {
					y++
				}
			}
			res += (x + y)
			n++
		}
	}
	glog.Info(n)
	return
}

func part2(lines []string) (res int) {
	galaxies := []util.Pos{}

	emptyRows := map[int]struct{}{}
	emptyCols := map[int]struct{}{}
	seenRows := map[int]struct{}{}
	seenCols := map[int]struct{}{}
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] != '.' {
				seenRows[r] = struct{}{}
				seenCols[c] = struct{}{}
				galaxies = append(galaxies, util.Pos{X: c, Y: r})
			}
		}
	}
	for r := 0; r < len(lines); r++ {
		if _, ok := seenRows[r]; !ok {
			emptyRows[r] = struct{}{}
		}
	}
	for c := 0; c < len(lines[0]); c++ {
		if _, ok := seenCols[c]; !ok {
			emptyCols[c] = struct{}{}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			xs := []int{g1.X, g2.X}
			slices.Sort(xs)
			x := xs[1] - xs[0]
			for c := range emptyCols {
				if c > xs[0] && c < xs[1] {
					x += 999999
				}
			}
			ys := []int{g1.Y, g2.Y}
			slices.Sort(ys)
			y := ys[1] - ys[0]
			for r := range emptyRows {
				if r > ys[0] && r < ys[1] {
					y += 999999
				}
			}
			res += (x + y)
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
