package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
)

func parseView(m []string, x, y, dx, dy int) (score int, visible bool) {
	visible = true
	n := len(m)
	i := x + dx
	j := y + dy
	for {
		score++
		if m[i][j] >= m[x][y] {
			visible = false
			return
		}
		i += dx
		j += dy
		if i < 0 || i == n || j < 0 || j == n {
			return
		}
	}
}

func parseTree(m []string, x, y int) (score int, visible bool) {
	st, vt := parseView(m, x, y, -1, 0)
	sb, vb := parseView(m, x, y, 1, 0)
	sl, vl := parseView(m, x, y, 0, -1)
	sr, vr := parseView(m, x, y, 0, 1)

	score = st * sb * sl * sr
	visible = vt || vb || vl || vr
	return
}

func isEdge(i, j, n int) bool {
	return i == 0 || j == 0 || i == n-1 || j == n-1
}

func survey(lines []string) (part1, part2 int) {
	n := len(lines)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isEdge(i, j, n) {
				part1++
				continue
			}
			score, visible := parseTree(lines, i, j)
			if visible {
				part1++
			}
			if score > part2 {
				part2 = score
			}
		}
	}
	return
}

func main() {
	var (
		t time.Time
	)
	lines := strings.Split(input, "\n")

	t = time.Now()
	part1, part2 := survey(lines)
	fmt.Printf("[Part 1] = %v\n", part1)
	fmt.Printf("[Part 2] = %v\n", part2)
	fmt.Printf("took %s\n\n", time.Since(t))
}
