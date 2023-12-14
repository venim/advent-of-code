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

func north(dish [][]rune) {
	for c := 0; c < len(dish[0]); c++ {
	row:
		for r := 0; r < len(dish); r++ {
			if dish[r][c] == 'O' {
				for r2 := r - 1; r2 >= -1; r2-- {
					if r2 == -1 || dish[r2][c] == 'O' || dish[r2][c] == '#' {
						dish[r][c] = '.'
						dish[r2+1][c] = 'O'
						continue row
					}
				}
			}
		}
	}
}

func west(dish [][]rune) {
	for r := 0; r < len(dish); r++ {
	col:
		for c := 0; c < len(dish[0]); c++ {
			if dish[r][c] == 'O' {
				for c2 := c - 1; c2 >= -1; c2-- {
					if c2 == -1 || dish[r][c2] == 'O' || dish[r][c2] == '#' {
						dish[r][c] = '.'
						dish[r][c2+1] = 'O'
						continue col
					}
				}
			}
		}
	}
}

func south(dish [][]rune) {
	for c := 0; c < len(dish[0]); c++ {
	row:
		for r := len(dish) - 1; r >= 0; r-- {
			if dish[r][c] == 'O' {
				for r2 := r + 1; r2 <= len(dish); r2++ {
					if r2 == len(dish) || dish[r2][c] == 'O' || dish[r2][c] == '#' {
						dish[r][c] = '.'
						dish[r2-1][c] = 'O'
						continue row
					}
				}
			}
		}
	}
}

func east(dish [][]rune) {
	for r := 0; r < len(dish); r++ {
	col:
		for c := len(dish[0]) - 1; c >= 0; c-- {
			if dish[r][c] == 'O' {
				for c2 := c + 1; c2 <= len(dish[0]); c2++ {
					if c2 == len(dish) || dish[r][c2] == 'O' || dish[r][c2] == '#' {
						dish[r][c] = '.'
						dish[r][c2-1] = 'O'
						continue col
					}
				}
			}
		}
	}
}

func score(dish [][]rune) (res int) {
	for c := 0; c < len(dish[0]); c++ {
		for r := 0; r < len(dish); r++ {
			if dish[r][c] == 'O' {
				res += (len(dish) - r)
			}
		}
	}
	return
}

func part1(lines []string) (res int) {
	dish := make([][]rune, len(lines))
	for i, l := range lines {
		dish[i] = []rune(l)
	}
	north(dish)
	return score(dish)
}

func toString(dish [][]rune) string {
	s := ""
	for _, r := range dish {
		s += string(r) + "\n"
	}
	return s[:len(s)-1]
}

func fromString(s string) [][]rune {
	lines := strings.Split(s, "\n")
	dish := make([][]rune, len(lines))
	for i, l := range lines {
		dish[i] = []rune(l)
	}
	return dish
}

func part2(lines []string) (res int) {
	dish := make([][]rune, len(lines))
	for i, l := range lines {
		dish[i] = []rune(l)
	}
	cache := []string{}

	for cycle := 0; cycle < 1000000000; cycle++ {
		north(dish)
		west(dish)
		south(dish)
		east(dish)
		if prev := slices.Index(cache, toString(dish)); prev != -1 {
			n := (1000000000 - (prev + 1)) % (cycle - prev)
			dish = fromString(cache[prev+n])
			break
		}
		cache = append(cache, toString(dish))
	}
	return score(dish)
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
