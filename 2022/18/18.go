package main

import (
	_ "embed"
	"flag"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type rock [3]int

type grid map[rock]int

func surfaceArea(r rock, rocks grid) int {
	sa := 0
	for i := 0; i < 3; i++ {
		for j := -1; j <= 1; j += 2 {
			adj := r
			adj[i] += j
			if !exists(adj, rocks) {
				sa++
			}
		}
	}
	return sa
}

func insideCube(r, max, min rock) bool {
	for i := 0; i < 3; i++ {
		if r[i] < min[i] {
			return false
		}
		if r[i] > max[i] {
			return false
		}
	}
	return true
}

func extSurfaceArea(rocks grid, max, min rock) (sa int) {
	q := []rock{min}
	seen := make(grid)
	for len(q) > 0 {
		r := q[0]
		q = q[1:]
		for i := 0; i < 3; i++ {
			for j := -1; j <= 1; j += 2 {
				adj := r
				adj[i] += j
				if exists(adj, rocks) {
					sa++
				} else if !exists(adj, seen) && insideCube(adj, max, min) {
					seen[adj] = 0
					q = append(q, adj)
				}
			}
		}
	}
	return
}

func exists(r rock, rocks grid) bool {
	_, ok := rocks[r]
	return ok
}

func makeRock(line string) rock {
	l := strings.Split(line, ",")
	return rock{mustAtoi(l[0]), mustAtoi(l[1]), mustAtoi(l[2])}
}

func part1(lines []string) (res int) {
	rocks := make(grid)
	for _, line := range lines {
		rocks[makeRock(line)] = 0
	}
	for rock := range rocks {
		res += surfaceArea(rock, rocks)
	}
	return
}

func part2(lines []string) (res int) {
	rocks := make(grid)
	max := rock{math.MinInt, math.MinInt, math.MinInt}
	min := rock{math.MaxInt, math.MaxInt, math.MaxInt}
	for _, line := range lines {
		r := makeRock(line)
		for i := 0; i < 3; i++ {
			if r[i] > max[i] {
				max[i] = r[i]
			}
			if r[i] < min[i] {
				min[i] = r[i]
			}
		}
		rocks[r] = 0
	}
	for i := 0; i < 3; i++ {
		max[i]++
		min[i]--
	}
	res += extSurfaceArea(rocks, max, min)
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
