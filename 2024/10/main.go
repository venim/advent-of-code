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

type trailMap struct {
	size  util.Pos
	peaks map[util.Pos]bool
	m     map[util.Pos]byte
}

func parse(lines []string) trailMap {
	m := trailMap{
		size:  util.Pos{len(lines[0]), len(lines)},
		m:     make(map[util.Pos]byte),
		peaks: make(map[util.Pos]bool),
	}
	var pos util.Pos
	for r := 0; r < m.size.Y; r++ {
		for c := 0; c < m.size.X; c++ {
			pos = util.Pos{c, r}
			m.m[pos] = lines[r][c]
			if lines[r][c] == '9' {
				m.peaks[pos] = true
			}
		}
	}
	return m
}

func findScore(m trailMap, pos util.Pos, peaks map[util.Pos]bool) (res int) {
	if m.m[pos] == '9' {
		if peaks == nil {
			return 1
		}
		if peaks[pos] {
			return 0
		}
		peaks[pos] = true
		return 1
	}
	for _, d := range util.Directions {
		p := pos.Add(d)
		if !p.IsOutOfBounds(m.size.X, m.size.Y) && m.m[p] == m.m[pos]+1 {
			res += findScore(m, p, peaks)
		}
	}
	return
}

func part1(lines []string) (res int) {
	m := parse(lines)
	for p, v := range m.m {
		if v == '0' {
			res += findScore(m, p, map[util.Pos]bool{})
		}
	}
	return
}

func part2(lines []string) (res int) {
	m := parse(lines)
	for p, v := range m.m {
		if v == '0' {
			res += findScore(m, p, nil)
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
