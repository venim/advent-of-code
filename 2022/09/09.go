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

type pos struct {
	X int
	Y int
}

type knot struct {
	Pos    pos
	Places map[pos]struct{}
}

type rope struct {
	Knots []knot
}

func newRope(n int) rope {
	r := rope{
		make([]knot, n+1),
	}
	for i := 0; i < n+1; i++ {
		r.Knots[i] = knot{Places: map[pos]struct{}{}}
	}
	return r
}

func (r *rope) moveHead(dir string) {
	switch dir {
	case "R":
		r.Knots[0].Pos.X++
	case "L":
		r.Knots[0].Pos.X--
	case "U":
		r.Knots[0].Pos.Y++
	case "D":
		r.Knots[0].Pos.Y--
	}
	r.moveTails()
}

func (r *rope) moveTails() {
	for i := 1; i < len(r.Knots); i++ {
		var head pos
		if i == 0 {
			continue
		} else {
			head = r.Knots[i-1].Pos
		}
		// newPos :=
		r.Knots[i].Pos = moveTailPos(head, r.Knots[i].Pos)
		r.Knots[i].Places[r.Knots[i].Pos] = struct{}{}
	}
}

func moveTailPos(head pos, tail pos) pos {
	dx := head.X - tail.X
	dy := head.Y - tail.Y
	switch dx {
	case 0:
		switch dy {
		case 2:
			tail.Y++
		case -2:
			tail.Y--
		}
	case 1:
		switch dy {
		case 2:
			tail.X++
			tail.Y++
		case -2:
			tail.X++
			tail.Y--
		}
	case 2:
		switch dy {
		case 0:
			tail.X++
		case 1, 2:
			tail.X++
			tail.Y++
		case -1, -2:
			tail.X++
			tail.Y--
		}
	case -1:
		switch dy {
		case 2:
			tail.X--
			tail.Y++
		case -2:
			tail.X--
			tail.Y--
		}
	case -2:
		switch dy {
		case 0:
			tail.X--
		case 1, 2:
			tail.X--
			tail.Y++
		case -1, -2:
			tail.X--
			tail.Y--
		}
	}
	return tail
}

// func (r *rope) draw() {
// 	n := 10
// 	lines := make([]string, n*2+1)
// 	for i := -n*2 + 1; i < 1; i++ {
// 		line := make([]rune, n*2+1)
// 		for j := -n; j <= n; j++ {
// 			line[j+n] = '.'
// 			if _, ok := r.Knots[1].Places[pos{i, j}]; ok {
// 				line[j+n] = '#'
// 			}
// 			if i == 0 && j == 0 {
// 				line[j+n] = 's'
// 			}
// 			if r.Knots[1].Pos.X == i && r.Knots[1].Pos.Y == j {
// 				line[j+n] = 'T'
// 			}
// 			if r.Knots[0].Pos.X == i && r.Knots[0].Pos.Y == j {
// 				line[j+n] = 'H'
// 			}
// 		}
// 		lines[i+n*2] = string(line)

// 	}
// 	fmt.Printf("%s\n\n", strings.Join(lines, "\n"))
// 	time.Sleep(10 * time.Millisecond)
// }

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part1(lines []string) (res int) {
	r := newRope(1)
	for _, line := range lines {
		l := strings.Split(line, " ")
		for i := 0; i < mustAtoi(l[1]); i++ {
			r.moveHead(l[0])
		}
	}
	return len(r.Knots[1].Places)
}

func part2(lines []string) (res int) {
	r := newRope(9)
	for _, line := range lines {
		l := strings.Split(line, " ")
		for i := 0; i < mustAtoi(l[1]); i++ {
			r.moveHead(l[0])
		}
	}
	return len(r.Knots[9].Places)
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
