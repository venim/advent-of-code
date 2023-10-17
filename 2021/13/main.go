package main

import (
	_ "embed"
	"flag"
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

type pos struct {
	X, Y int
}

type board [][]bool

type fold struct {
	dir byte
	n   int
}

func newBoard(x, y int) board {
	b := make(board, x)
	for i := 0; i < x; i++ {
		b[i] = make([]bool, y)
		for j := 0; j < y; j++ {
			b[i][j] = false
		}
	}
	return b
}

func (b board) countDots() int {
	dots := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] {
				dots++
			}
		}
	}
	return dots
}

func (b board) draw() {
	for j := 0; j < len(b[0]); j++ {
		var s string
		for i := 0; i < len(b); i++ {
			if b[i][j] {
				s += "#"
			} else {
				s += "."
			}
		}
		glog.Info(s)
	}
}

func (b board) fold(f fold) board {
	switch f.dir {
	case 'x':
		return b.foldX(f.n)
	case 'y':
		return b.foldY(f.n)
	}
	return b
}

func reorder[T any](x []T, y []T) ([]T, []T) {
	if len(x) < len(y) {
		return y, x
	}
	return x, y
}

func (b board) foldX(n int) board {
	newX := max(len(b)-n-1, n)
	newB := newBoard(newX, len(b[0]))

	left := b[:n]
	right := b[n+1:]

	// reverse the right side
	for i, j := 0, len(right)-1; i < j; i, j = i+1, j-1 {
		right[i], right[j] = right[j], right[i]
	}

	// ensure left is the bigger array
	left, right = reorder(left, right)

	// pad the smaller array
	if len(right) < newX {
		tmp := newBoard(newX, len(b[0]))
		copy(tmp[newX-len(right):], right[:])
		right = tmp
	}

	for i := 0; i < len(newB); i++ {
		for j := 0; j < len(b[0]); j++ {
			newB[i][j] = left[i][j] || right[i][j]
		}
	}

	return newB
}

func (b board) foldY(n int) board {
	newY := max(len(b[0])-n-1, n)
	newB := newBoard(len(b), newY)

	for i := 0; i < len(newB); i++ {
		top := b[i][:n]
		bot := b[i][n+1:]
		// reverse bottom
		for i2, j := 0, len(bot)-1; i2 < j; i2, j = i2+1, j-1 {
			bot[i2], bot[j] = bot[j], bot[i2]
		}

		// make sure top is the larger array
		top, bot = reorder(top, bot)

		// left pad the smaller array
		if len(bot) < newY {
			tmp := make([]bool, newY)
			copy(tmp[newY-len(bot):], bot[:])
			bot = tmp
		}

		for j := 0; j < newY; j++ {
			newB[i][j] = top[j] || bot[j]
		}
	}
	return newB
}

func parse(lines []string) (board, []fold) {
	dotsStr, foldsStr, _ := strings.Cut(strings.Join(lines, "\n"), "\n\n")
	dotLines := strings.Split(dotsStr, "\n")
	foldLines := strings.Split(foldsStr, "\n")
	maxX := -1
	maxY := -1

	dots := make([]pos, len(dotLines))
	for i, d := range dotLines {
		xStr, yStr, _ := strings.Cut(d, ",")
		x := mustAtoi(xStr)
		maxX = max(maxX, x)
		y := mustAtoi(yStr)
		maxY = max(maxY, y)
		dots[i] = pos{x, y}
	}

	b := newBoard(maxX+1, maxY+1)

	for _, d := range dots {
		b[d.X][d.Y] = true
	}

	folds := make([]fold, 0, len(foldLines))
	for _, f := range foldLines {
		foldAxis, foldN, _ := strings.Cut(f, "=")
		n := mustAtoi(foldN)
		dir := foldAxis[len(foldAxis)-1]
		folds = append(folds, fold{dir, n})
	}
	return b, folds
}

func part1(lines []string) (res int) {
	board, folds := parse(lines)
	board = board.fold(folds[0])

	return board.countDots()
}

func part2(lines []string) (res int) {
	board, folds := parse(lines)
	for _, f := range folds {
		board = board.fold(f)
	}

	board.draw()
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
