package main

import (
	_ "embed"
	"flag"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

type dish [][]rune

func (d dish) tilt() {
	for c := 0; c < len(d[0]); c++ {
	row:
		for r := 0; r < len(d); r++ {
			if d[r][c] == 'O' {
				for rr := r - 1; rr >= -1; rr-- {
					if rr == -1 || d[rr][c] == 'O' || d[rr][c] == '#' {
						d[r][c] = '.'
						d[rr+1][c] = 'O'
						continue row
					}
				}
			}
		}
	}
}

func (d dish) calcLoad() (res int) {
	for c := 0; c < len(d[0]); c++ {
		for r := 0; r < len(d); r++ {
			if d[r][c] == 'O' {
				res += (len(d) - r)
			}
		}
	}
	return
}

func part1(lines []string) (res int) {
	dish := dishFromString(strings.Join(lines, "\n"))
	dish.tilt()
	return dish.calcLoad()
}

func (d *dish) rotate() {
	n := len(*d)
	m := len((*d)[0])
	dd := make(dish, m)
	for j := 0; j < m; j++ {
		dd[j] = make([]rune, n)
		for i := 0; i < n; i++ {
			dd[j][n-1-i] = (*d)[i][j]
		}
	}
	*d = dd
}

func (d dish) String() string {
	s := ""
	for _, r := range d {
		s += string(r) + "\n"
	}
	return s[:len(s)-1]
}

func dishFromString(s string) dish {
	lines := strings.Split(s, "\n")
	d := make(dish, len(lines))
	for i, l := range lines {
		d[i] = []rune(l)
	}
	return d
}

func (d *dish) cycle(n int) {
	cache := []string{}
	for cycle := 0; cycle < n; cycle++ {
		for i := 0; i < 4; i++ {
			d.tilt()
			d.rotate()
		}
		s := d.String()
		if prev := slices.Index(cache, s); prev != -1 {
			n := (999999999 - prev) % (cycle - prev)
			*d = dishFromString(cache[prev+n])
			return
		}
		cache = append(cache, s)
	}
}

func part2(lines []string) (res int) {
	dish := dishFromString(strings.Join(lines, "\n"))
	dish.cycle(1000000000)
	return dish.calcLoad()
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
