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

func (d dish) rollNorth() {
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

func (d dish) rollWest() {
	for r := 0; r < len(d); r++ {
	col:
		for c := 0; c < len(d[0]); c++ {
			if d[r][c] == 'O' {
				for cc := c - 1; cc >= -1; cc-- {
					if cc == -1 || d[r][cc] == 'O' || d[r][cc] == '#' {
						d[r][c] = '.'
						d[r][cc+1] = 'O'
						continue col
					}
				}
			}
		}
	}
}

func (d dish) rollSouth() {
	for c := 0; c < len(d[0]); c++ {
	row:
		for r := len(d) - 1; r >= 0; r-- {
			if d[r][c] == 'O' {
				for rr := r + 1; rr <= len(d); rr++ {
					if rr == len(d) || d[rr][c] == 'O' || d[rr][c] == '#' {
						d[r][c] = '.'
						d[rr-1][c] = 'O'
						continue row
					}
				}
			}
		}
	}
}

func (d dish) rollEast() {
	for r := 0; r < len(d); r++ {
	col:
		for c := len(d[0]) - 1; c >= 0; c-- {
			if d[r][c] == 'O' {
				for cc := c + 1; cc <= len(d[0]); cc++ {
					if cc == len(d) || d[r][cc] == 'O' || d[r][cc] == '#' {
						d[r][c] = '.'
						d[r][cc-1] = 'O'
						continue col
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
	dish.rollNorth()
	return dish.calcLoad()
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

func part2(lines []string) (res int) {
	dish := dishFromString(strings.Join(lines, "\n"))
	cache := []string{}

	for cycle := 0; cycle < 1000000000; cycle++ {
		dish.rollNorth()
		dish.rollWest()
		dish.rollSouth()
		dish.rollEast()
		if prev := slices.Index(cache, dish.String()); prev != -1 {
			n := (999999999 - prev) % (cycle - prev)
			dish = dishFromString(cache[prev+n])
			break
		}
		cache = append(cache, dish.String())
	}
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
