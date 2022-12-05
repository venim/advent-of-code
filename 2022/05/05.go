package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
	lines []string
)

func parseBoard(board []string) (int, [][]string) {
	bottom := -1
	for i, line := range board {
		if line == "" {
			bottom = i
			break
		}
	}
	out := make([][]string, len(board[bottom-1])/4+1)
	for i := bottom - 2; i >= 0; i-- {
		for j := 0; j < len(board[i])-1; j += 4 {
			if board[i][j+1] != ' ' {
				out[j/4] = append(out[j/4], string(board[i][j+1]))
			}
		}
	}

	return bottom, out
}

func parseMove(line string) (n, from, to int) {
	atoi := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	parts := strings.Split(line, " ")
	return atoi(parts[1]), atoi(parts[3]), atoi(parts[5])
}

func pop[T any](s []T) (T, []T) {
	if n := len(s); n > 0 {
		return s[n-1], s[:n-1]
	}
	return *new(T), []T{}
}

type board struct {
	b [][]string
}

func (b board) move(n, from, to int) {
	if l := len(b.b[from-1]); n <= l {
		v := b.b[from-1][l-n:]
		// reverse slice
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
		b.b[to-1] = append(b.b[to-1], v...)
		b.b[from-1] = b.b[from-1][:l-n]
	}
}

func (b board) moveMultiple(n, from, to int) {
	if l := len(b.b[from-1]); n <= l {
		v := b.b[from-1][l-n:]
		b.b[to-1] = append(b.b[to-1], v...)
		b.b[from-1] = b.b[from-1][:l-n]
	}
}

func (b board) makeResult() (res string) {
	for _, s := range b.b {
		v, _ := pop(s)

		if v != "" {
			res += v
		} else {
			res += " "
		}
	}
	return
}

func part1(lines []string) string {
	i, b := parseBoard(lines)
	stacks := board{b}
	for _, line := range lines[i+1:] {
		if line == "" {
			continue
		}
		stacks.move(parseMove(line))
	}
	return stacks.makeResult()
}

func part2(lines []string) string {
	i, b := parseBoard(lines)
	stacks := board{b}
	for _, line := range lines[i+1:] {
		if line == "" {
			continue
		}
		stacks.moveMultiple(parseMove(line))
	}
	return stacks.makeResult()
}

func main() {
	lines = strings.Split(input, "\n")
	var (
		t   time.Time
		res string
	)

	t = time.Now()
	res = part1(lines)
	fmt.Printf("[Part 1] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	res = part2(lines)
	fmt.Printf("[Part 2] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))
}
