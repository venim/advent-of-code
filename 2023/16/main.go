package main

import (
	_ "embed"
	"flag"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

type grid map[complex128]byte

func makeGrid(lines []string) grid {
	g := grid{}
	for r := 0; r < len(lines[0]); r++ {
		for c := 0; c < len(lines); c++ {
			g[complex(float64(c), float64(r))] = lines[r][c]
		}
	}
	return g
}

type state struct {
	pos complex128
	dir complex128
}

func (G grid) solve(queue []state) int {
	done := map[state]bool{}
	ps := map[complex128]bool{}
	for len(queue) > 0 {
		pos := queue[0].pos
		dir := queue[0].dir
		queue = queue[1:]

		for !done[state{pos, dir}] {
			done[state{pos, dir}] = true
			ps[pos] = true
			pos += dir
			v, ok := G[pos]
			if !ok {
				break
			}
			switch v {
			case '|':
				dir = 1i
				queue = append(queue, state{pos, -dir})
			case '-':
				dir = 1
				queue = append(queue, state{pos, -dir})
			case '/':
				dir = -complex(imag(dir), real(dir))
			case '\\':
				dir = complex(imag(dir), real(dir))
			}
		}
	}

	return len(ps) - 1
}

func part1(lines []string) (res int) {
	g := makeGrid(lines)
	return g.solve([]state{{-1, 1}})
}

func part2(lines []string) (res int) {
	g := makeGrid(lines)
	for p := range g {
		for _, d := range []complex128{1, -1, 1i, -1i} {
			if _, ok := g[p-d]; !ok {
				res = max(res, g.solve([]state{{p - d, d}}))
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
