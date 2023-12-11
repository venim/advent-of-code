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

type set map[int]struct{}

type image struct {
	galaxies  []util.Pos
	emptyRows set
	emptyCols set
	emptyMult int
}

func findEmpty(n int, seen set) set {
	empty := set{}
	for i := 0; i < n; i++ {
		if _, ok := seen[i]; !ok {
			empty[i] = struct{}{}
		}
	}
	return empty
}

func parse(lines []string) image {
	galaxies := []util.Pos{}

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] != '.' {
				galaxies = append(galaxies, util.Pos{X: c, Y: r})
			}
		}
	}

	seenRows := set{}
	seenCols := set{}

	for _, galaxy := range galaxies {
		seenRows[galaxy.Y] = struct{}{}
		seenCols[galaxy.X] = struct{}{}
	}

	return image{
		galaxies:  galaxies,
		emptyRows: findEmpty(len(lines), seenRows),
		emptyCols: findEmpty(len(lines[0]), seenCols),
	}
}

func dist(vs []int, empties set, n int) int {
	slices.Sort(vs)
	v := vs[1] - vs[0]
	for i := range empties {
		if i > vs[0] && i < vs[1] {
			v += (n - 1)
		}
	}
	return v
}

func (img image) sumAllDistances() (res int) {
	for i := 0; i < len(img.galaxies); i++ {
		for j := i + 1; j < len(img.galaxies); j++ {
			g1 := img.galaxies[i]
			g2 := img.galaxies[j]

			res += dist([]int{g1.X, g2.X}, img.emptyCols, img.emptyMult)
			res += dist([]int{g1.Y, g2.Y}, img.emptyRows, img.emptyMult)
		}
	}
	return
}

func part1(lines []string) (res int) {
	img := parse(lines)
	img.emptyMult = 2
	return img.sumAllDistances()
}

func part2(lines []string) (res int) {
	img := parse(lines)
	img.emptyMult = 1000000
	return img.sumAllDistances()
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
