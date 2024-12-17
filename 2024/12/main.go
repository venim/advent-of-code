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

type garden struct {
	plots map[util.Pos]byte
	size  util.Pos
}

type region struct {
	plots map[util.Pos]bool
}

func (r region) findPerimeter() (perimeter int) {
	for p := range r.plots {
		perimeter += 4
		for _, d := range util.Directions {
			if r.plots[p.Add(d)] {
				perimeter--
			}
		}
	}
	return
}

func (r region) findNumSides() (sides int) {
	for p := range r.plots {
		sides += r.findCornersAtPos(p)
	}
	return
}

func (r region) findCornersAtPos(p util.Pos) (corners int) {
	adjs := [][]util.Pos{
		{{X: 1, Y: 0}, {X: 0, Y: 1}},   // right / down
		{{X: -1, Y: 0}, {X: 0, Y: -1}}, // left / up
		{{X: 1, Y: 0}, {X: 0, Y: -1}},  // right / up
		{{X: -1, Y: 0}, {X: 0, Y: 1}},  // left / down
	}
	for _, adjs := range adjs {
		edges := 0
		for _, d := range adjs {
			next := p.Add(d)
			if !r.plots[next] {
				edges++
			}
		}
		if edges == 2 {
			corners++
		} else if edges == 0 {
			diag := p.Add(adjs[0]).Add(adjs[1])
			if !r.plots[diag] {
				corners++
			}
		}
	}
	return
}

func parse(lines []string) *garden {
	plots := map[util.Pos]byte{}
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			plots[util.Pos{X: c, Y: r}] = lines[r][c]
		}
	}
	return &garden{
		plots: plots,
		size:  util.Pos{X: len(lines[0]), Y: len(lines)},
	}
}

func flood(g *garden, start util.Pos) *region {
	plots := map[util.Pos]bool{start: true}

	q := []util.Pos{start}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, d := range util.Directions {
			next := cur.Add(d)
			if _, ok := plots[next]; ok {
				// skip if we've checked this plot already
				continue
			}
			if !next.IsOutOfBounds(g.size.X, g.size.Y) && g.plots[next] == g.plots[start] {
				plots[next] = true
				q = append(q, next)
			}
		}
	}

	return &region{
		plots: plots,
	}
}

func part1(lines []string) (res int) {
	g := parse(lines)
	inRegion := map[util.Pos]bool{}
	for i := 0; i < g.size.X; i++ {
		for j := 0; j < g.size.Y; j++ {
			p := util.Pos{X: i, Y: j}
			// is this pos already in a region?
			if !inRegion[p] {
				// no, start flood
				region := flood(g, p)
				for r := range region.plots {
					inRegion[r] = true
				}

				res += len(region.plots) * region.findPerimeter()
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	g := parse(lines)
	inRegion := map[util.Pos]bool{}
	for i := 0; i < g.size.X; i++ {
		for j := 0; j < g.size.Y; j++ {
			p := util.Pos{X: i, Y: j}
			// is this pos already in a region?
			if !inRegion[p] {
				// no, start flood
				region := flood(g, p)
				for r := range region.plots {
					inRegion[r] = true
				}

				sides := region.findNumSides()
				res += len(region.plots) * sides
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
