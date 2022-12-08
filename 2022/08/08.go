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
)

func isEdge(i, j, n int) bool {
	if i == 0 || j == 0 || i == n-1 || j == n-1 {
		return true
	}
	return false
}

func isTall(m []string, x, y int) bool {
	n := len(m)
	h := m[x][y]

	top := true
	for i := 0; i < x; i++ {
		if m[i][y] >= h {
			top = false
			break
		}
	}
	bottom := true
	for i := x + 1; i < n; i++ {
		if m[i][y] >= h {
			bottom = false
			break
		}
	}
	left := true
	for i := 0; i < y; i++ {
		if m[x][i] >= h {
			left = false
			break
		}
	}
	right := true
	for i := y + 1; i < n; i++ {
		if m[x][i] >= h {
			right = false
			break
		}
	}
	return top || bottom || left || right
}

func scenicScore(m []string, x, y int) int {
	n := len(m)
	h := m[x][y]

	top := 0
	for i := x - 1; i >= 0; i-- {
		top++
		if m[i][y] >= h {
			break
		}
	}
	bottom := 0
	for i := x + 1; i < n; i++ {
		bottom++
		if m[i][y] >= h {
			break
		}
	}
	left := 0
	for i := y - 1; i >= 0; i-- {
		left++
		if m[x][i] >= h {
			break
		}
	}
	right := 0
	for i := y + 1; i < n; i++ {
		right++
		if m[x][i] >= h {
			break
		}
	}
	return top * bottom * left * right
}

func isVisible(m []string) [][]bool {
	n := len(m)
	res := make([][]bool, n)
	for i := 0; i < n; i++ {
		res[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if isEdge(i, j, n) {
				res[i][j] = true
				continue
			}
			tall := isTall(m, i, j)
			if tall {
				res[i][j] = true
			}
		}
	}
	return res
}

func scoreTrees(m []string) [][]int {
	n := len(m)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isEdge(i, j, n) {
				res[i][j] = 0
				continue
			}
			res[i][j] = scenicScore(m, i, j)
		}
	}
	return res
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part1(lines []string) (res int) {
	m := isVisible(lines)
	for _, r := range m {
		for _, t := range r {
			if t {
				res++
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	m := scoreTrees(lines)
	for _, r := range m {
		for _, t := range r {
			if t > res {
				res = t
			}
		}
	}
	return
}

func main() {
	var (
		t   time.Time
		res int
	)
	lines := strings.Split(input, "\n")

	t = time.Now()
	res = part1(lines)
	fmt.Printf("[Part 1] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	res = part2(lines)
	fmt.Printf("[Part 2] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))
}
