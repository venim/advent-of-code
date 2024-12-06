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
	XMAS  = []byte{'X', 'M', 'A', 'S'}
)

func validCoordinate(r, c, x, y int) bool {
	return r >= 0 && c >= 0 && r < x && c < y
}

func findXMAS(lines []string, dx, dy, r, c, x, y, depth int) (res int) {
	if depth == 4 {
		return 1
	}
	i := r + dy
	j := c + dx
	if validCoordinate(i, j, x, y) {
		if lines[i][j] == XMAS[depth] {
			res += findXMAS(lines, dx, dy, i, j, x, y, depth+1)
		}
	}

	return
}

func part1(lines []string) (res int) {
	x := len(lines)
	y := len(lines[0])
	for r := 0; r < x; r++ {
		for c := 0; c < y; c++ {
			if lines[r][c] == XMAS[0] {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						res += findXMAS(lines, i, j, r, c, x, y, 1)
					}
				}
			}
		}
	}
	return
}

var MAS = map[byte]byte{'M': 'S', 'S': 'M'}

func findX_MAS(lines []string, r, c, x, y int) bool {
	for _, j := range []int{c - 1, c + 1} {
		i := r - 1
		if !validCoordinate(i, j, x, y) {
			return false
		}
		want, ok := MAS[lines[i][j]]
		if !ok {
			return false
		}
		i += 2
		j += 2 * (c-j)
		if !validCoordinate(i, j, x, y) {
			return false
		}
		if lines[i][j] != want {
			return false
		}
	}
	return true
}

func part2(lines []string) (res int) {
	x := len(lines)
	y := len(lines[0])
	for r := 0; r < x; r++ {
		for c := 0; c < y; c++ {
			if lines[r][c] == 'A' {
				if findX_MAS(lines, r, c, x, y) {
					res++
				}
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
