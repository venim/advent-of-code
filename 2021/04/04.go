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
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

type data struct {
	Numbers []int
	Boards  []board
}

type board struct {
	Spots   map[pos]bool
	Numbers map[int]pos
}

func (b board) won() bool {
	for i := 0; i < 5; i++ {
		rowWon := true
		colWon := true
		for j := 0; j < 5; j++ {
			if spot, ok := b.Spots[pos{i, j}]; !ok || !spot {
				colWon = false
			}
			if spot, ok := b.Spots[pos{j, i}]; !ok || !spot {
				rowWon = false
			}
		}
		if colWon {
			return true
		}
		if rowWon {
			return true
		}
	}
	return false
}

func (b board) sum() int {
	sum := 0
	for n, pos := range b.Numbers {
		if called, _ := b.Spots[pos]; !called {
			sum += n
		}
	}
	return sum
}

func (b board) callN(n int) bool {
	if pos, ok := b.Numbers[n]; ok {
		b.Spots[pos] = true
		if b.won() {
			return true
		}
	}
	return false
}

func parseInput(lines []string) *data {
	// first line are the numbers called
	ns := strings.Split(lines[0], ",")
	numbers := make([]int, 0, len(ns))
	for _, i := range ns {
		n := mustAtoi(i)
		numbers = append(numbers, n)
	}

	// parse each 5 row board (+ newline)
	boards := []board{}
	for i := 2; i < len(lines); i += 6 {
		board := board{
			Spots:   map[pos]bool{},
			Numbers: map[int]pos{},
		}
		for j := i; j < i+5; j++ {
			line := strings.Fields(lines[j])
			for k, n := range line {
				pos := pos{k, j - i}
				board.Numbers[mustAtoi(n)] = pos
			}
		}
		boards = append(boards, board)
	}
	return &data{numbers, boards}
}

func part1(lines []string) (res int) {
	data := parseInput(lines)
	for _, n := range data.Numbers {
		for _, b := range data.Boards {
			if won := b.callN(n); won {
				return n * b.sum()
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	data := parseInput(lines)
	for _, n := range data.Numbers {
		remaining := make([]board, 0, len(data.Boards))
		for _, b := range data.Boards {
			if won := b.callN(n); won {
				if len(data.Boards) > 1 {
					continue
				}
				return n * b.sum()
			}
			remaining = append(remaining, b)
		}
		data.Boards = remaining
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
