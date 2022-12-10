package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
)

type cpu struct {
	Cycle int
	X     int
}

func (c *cpu) tick(v int) {
	c.Cycle++
	if v != 0 {
		c.X += v
	}
}

func newCPU() *cpu {
	return &cpu{
		Cycle: 1,
		X:     1,
	}
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func run(lines []string, cycles []int) (part1 int, part2 string) {
	c := newCPU()
	sort.Ints(cycles)
	n := cycles[len(cycles)-1]
	wantCycles := map[int]struct{}{}
	for _, c := range cycles {
		wantCycles[c] = struct{}{}
	}
	ticks := make([]int, 0, n)

	for _, cmd := range lines {
		if strings.Index(cmd, "noop") == 0 {
			ticks = append(ticks, 0)
		} else {
			v := mustAtoi(strings.Split(cmd, " ")[1])
			ticks = append(ticks, 0)
			ticks = append(ticks, v)
		}
	}
	nPerLine := 40
	line := 0

	for i, tick := range ticks {
		// draw pixel
		pos := (c.Cycle - 1) - (nPerLine * line)
		if c.X-1 <= pos && pos <= c.X+1 {
			part2 += "#"
		} else {
			part2 += "."
		}
		if (i+1)%nPerLine == 0 {
			part2 += "\n"
			line++
		}
		// do tick
		c.tick(tick)
		// record strength
		if _, ok := wantCycles[i+2]; ok {
			part1 += c.Cycle * c.X
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
	part1, part2 := run(lines, []int{20, 60, 100, 140, 180, 220})
	fmt.Printf("[Part 1] = %v\n", part1)
	fmt.Printf("[Part 2] = \n%v\n", part2)
	fmt.Printf("took %s\n\n", time.Since(t))
}
