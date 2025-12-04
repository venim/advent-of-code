package main

import (
	_ "embed"
	"flag"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func sanitize(lines []string) [][]rune {
	board := [][]rune{}
	for _, l := range lines {
		board = append(board, []rune(l))
	}
	return board
}

func neighbors(lines [][]rune, pos util.Pos, x, y int) (res int) {
	for j := pos.Y - 1; j <= pos.Y+1; j++ {
		for i := pos.X - 1; i <= pos.X+1; i++ {
			if !(i == pos.X && j == pos.Y) && // don't check the center position
				!(util.Pos{i, j}).IsOutOfBounds(x, y) { // don't check if its not a position
				if lines[j][i] == '@' {
					res++
				}
			}
		}
	}
	return
}

func findAccessible(lines [][]rune) (accessible []util.Pos) {
	x := len(lines[0])
	y := len(lines)
	for j := range y {
		for i := range x {
			pos := util.Pos{i, j}
			if lines[j][i] == '@' &&
				neighbors(lines, pos, x, y) < 4 {
				accessible = append(accessible, pos)
			}
		}
	}
	return
}

func part1(lines [][]rune) (res int) {
	return len(findAccessible(lines))
}

func part2(lines [][]rune) (res int) {
	for {
		accessible := findAccessible(lines)
		if len(accessible) == 0 {
			return
		}
		res += len(accessible)
		// remove all accessible rolls
		for _, p := range accessible {
			lines[p.Y][p.X] = '.'
		}
	}
}

func init() {
	_ = flag.Set("logtostderr", "true")
}

func main() {
	var (
		t   time.Time
		res int
	)
	flag.Parse()
	lines := strings.Split(input, "\n")
	board := sanitize(lines)

	t = time.Now()
	res = part1(board)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(board)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
