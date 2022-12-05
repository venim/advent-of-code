package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-collections/collections/stack"
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

type board struct {
	stack []stack.Stack
}

func (b board) move(n, from, to int) {
	for i := 0; i < n; i++ {
		v := b.stack[from-1].Pop()
		if v == nil {
			continue
		}
		b.stack[to-1].Push(v)
	}
}

func (b board) moveMultiple(n, from, to int) {
	tmp := make([]any, 0, b.stack[from-1].Len())
	for i := 0; i < n; i++ {
		v := b.stack[from-1].Pop()
		if v == nil {
			continue
		}
		tmp = append(tmp, v)
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		b.stack[to-1].Push(tmp[i])
	}
}

func (b board) makeResult() (res string) {
	for _, s := range b.stack {
		v := s.Pop()

		if v != nil {
			res += v.(string)
		} else {
			res += " "
		}
	}
	return
}

func makeStacks(board [][]string) []stack.Stack {
	stacks := make([]stack.Stack, len(board))
	for i, row := range board {
		for _, v := range row {
			stacks[i].Push(v)
		}
	}
	return stacks
}

func part1(lines []string) string {
	i, b := parseBoard(lines)
	stacks := board{makeStacks(b)}
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
	stacks := board{makeStacks(b)}
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
