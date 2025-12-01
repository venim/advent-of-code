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

func part1(lines []string) (res int) {
	pos := 50
	for _, l := range lines {
		if l == "" {
			break
		}
		n := util.MustAtoi(l[1:])
		switch l[0] {
		case 'L':
			pos -= n
		case 'R':
			pos += n
		}
		// Ensure pos stays within [0, 99] using modulo arithmetic.
		// The `+ 100` ensures a positive result before the final modulo.
		pos = (pos%100 + 100) % 100
		if pos == 0 {
			res++
		}
	}
	return
}

func part2(lines []string) (res int) {
	// Initialize pos with a large offset to prevent negative values
	// and simplify boundary crossing calculations with integer division.
	pos := 50 + 1_000_000_000
	for _, l := range lines {
		if l == "" {
			break
		}
		n := util.MustAtoi(l[1:])
		switch l[0] {
		case 'L':
			next := pos - n
			// Count how many multiples of 100 are crossed when moving left.
			// (pos-1)/100 gives the number of full 100-segments before current pos.
			// (next-1)/100 gives the number of full 100-segments before next pos.
			// The difference is the number of boundaries crossed.
			res += (pos-1)/100 - (next-1)/100
			pos = next
		case 'R':
			next := pos + n
			// Count how many multiples of 100 are crossed when moving right.
			// next/100 gives the number of full 100-segments up to next pos.
			// pos/100 gives the number of full 100-segments up to current pos.
			// The difference is the number of boundaries crossed.
			res += next/100 - pos/100
			pos = next
		}
	}
	return
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

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
