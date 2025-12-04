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

func initGraph(lines [][]rune) ([]util.Pos, map[util.Pos]int) {
	x := len(lines[0])
	y := len(lines)
	counts := map[util.Pos]int{}
	q := []util.Pos{}

	for j := range y {
		for i := range x {
			if lines[j][i] == '@' {
				pos := util.Pos{i, j}
				c := neighbors(lines, pos, x, y)
				counts[pos] = c
				if c < 4 {
					q = append(q, pos)
				}
			}
		}
	}
	return q, counts
}

func part1(lines [][]rune) (res int) {
	q, _ := initGraph(lines)
	return len(q)
}

func part2(lines [][]rune) (res int) {
	x := len(lines[0])
	y := len(lines)
	q, counts := initGraph(lines)

	for head := 0; head < len(q); head++ {
		curr := q[head]
		lines[curr.Y][curr.X] = '.'

		for j := curr.Y - 1; j <= curr.Y+1; j++ {
			for i := curr.X - 1; i <= curr.X+1; i++ {
				if i == curr.X && j == curr.Y {
					continue
				}
				p := util.Pos{i, j}
				if !p.IsOutOfBounds(x, y) && lines[j][i] == '@' {
					counts[p]--
					if counts[p] == 3 {
						q = append(q, p)
					}
				}
			}
		}
	}
	return len(q)
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
