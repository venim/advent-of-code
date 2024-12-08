package main

import (
	_ "embed"
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

type grid struct {
	bounds util.Pos
	nodes  map[byte][]util.Pos
}

func parse(lines []string) grid {
	bounds := util.Pos{X: len(lines), Y: len(lines[0])}
	nodes := map[byte][]util.Pos{}
	for r := 0; r < bounds.X; r++ {
		for c := 0; c < bounds.Y; c++ {
			if v := lines[r][c]; v != '.' {
				p := util.Pos{X: r, Y: c}
				if _, ok := nodes[v]; !ok {
					nodes[v] = []util.Pos{}
				}
				nodes[v] = append(nodes[v], p)
			}
		}
	}
	return grid{
		bounds: bounds,
		nodes:  nodes,
	}
}

func extendLine(a, b, bounds util.Pos, n int) []util.Pos {
	p := b
	d := b.Sub(a)
	i := 0
	ans := []util.Pos{}
	for {
		p = p.Add(d)
		if p.IsOutOfBounds(bounds.X, bounds.Y) || (n != -1 && i >= n) {
			break
		}
		i++
		ans = append(ans, p)
	}
	return ans
}

func findAntinodes(a, b, bounds util.Pos, n int) []util.Pos {
	ans := []util.Pos{}
	ans = append(ans, extendLine(a, b, bounds, n)...)
	ans = append(ans, extendLine(b, a, bounds, n)...)
	return ans
}

func part1(lines []string) (res int) {
	grid := parse(lines)

	anodes := map[util.Pos]bool{}
	for _, ns := range grid.nodes {
		for i := 0; i < len(ns); i++ {
			for j := i + 1; j < len(ns); j++ {
				for _, an := range findAntinodes(ns[i], ns[j], grid.bounds, 1) {
					anodes[an] = true
				}
			}
		}
	}
	return len(anodes)
}

func part2(lines []string) (res int) {
	grid := parse(lines)

	anodes := map[util.Pos]bool{}
	for _, ns := range grid.nodes {
		for i := 0; i < len(ns); i++ {
			anodes[ns[i]] = true
			for j := i + 1; j < len(ns); j++ {
				for _, an := range findAntinodes(ns[i], ns[j], grid.bounds, -1) {
					anodes[an] = true
				}
			}
		}
	}
	return len(anodes)
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
