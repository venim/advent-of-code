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

func north(lines []string) {
	for c := 0; c < len(lines[0]); c++ {
	row:
		for r := 0; r < len(lines); r++ {
			if lines[r][c] == 'O' {
				for r2 := r - 1; r2 >= -1; r2-- {
					if r2 == -1 || lines[r2][c] == 'O' || lines[r2][c] == '#' {
						l := []rune(lines[r])
						l[c] = '.'
						lines[r] = string(l)
						l2 := []rune(lines[r2+1])
						l2[c] = 'O'
						lines[r2+1] = string(l2)
						continue row
					}
				}
			}
		}
	}
}

func west(lines []string) {
	for r := 0; r < len(lines); r++ {
	col:
		for c := 0; c < len(lines[0]); c++ {
			if lines[r][c] == 'O' {
				for c2 := c - 1; c2 >= -1; c2-- {
					if c2 == -1 || lines[r][c2] == 'O' || lines[r][c2] == '#' {
						l := []rune(lines[r])
						l[c] = '.'
						l[c2+1] = 'O'
						lines[r] = string(l)
						continue col
					}
				}
			}
		}
	}
}

func south(lines []string) {
	for c := 0; c < len(lines[0]); c++ {
	row:
		for r := len(lines) - 1; r >= 0; r-- {
			if lines[r][c] == 'O' {
				for r2 := r + 1; r2 <= len(lines); r2++ {
					if r2 == len(lines) || lines[r2][c] == 'O' || lines[r2][c] == '#' {
						l := []rune(lines[r])
						l[c] = '.'
						lines[r] = string(l)
						l2 := []rune(lines[r2-1])
						l2[c] = 'O'
						lines[r2-1] = string(l2)
						continue row
					}
				}
			}
		}
	}
}

func east(lines []string) {
	for r := 0; r < len(lines); r++ {
	col:
		for c := len(lines[0]) - 1; c >= 0; c-- {
			if lines[r][c] == 'O' {
				for c2 := c + 1; c2 <= len(lines[0]); c2++ {
					if c2 == len(lines) || lines[r][c2] == 'O' || lines[r][c2] == '#' {
						l := []rune(lines[r])
						l[c] = '.'
						l[c2-1] = 'O'
						lines[r] = string(l)
						continue col
					}
				}
			}
		}
	}
}

func score(lines []string) (res int) {
	for c := 0; c < len(lines[0]); c++ {
		for r := 0; r < len(lines); r++ {
			if lines[r][c] == 'O' {
				res += (len(lines) - r)
			}
		}
	}
	return
}

func part1(lines []string) (res int) {
	north(lines)
	return score(lines)
}

func part2(lines []string) (res int) {
	cache := []string{}

	for cycle := 0; cycle < 1000000000; cycle++ {
		north(lines)
		west(lines)
		south(lines)
		east(lines)
		if prev := slices.Index(cache, strings.Join(lines, "\n")); prev != -1 {
			n := (1000000000 - (prev + 1)) % (cycle - prev)
			lines = strings.Split(cache[prev+n], "\n")
			break
		}
		cache = append(cache, strings.Join(lines, "\n"))
	}
	return score(lines)
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
