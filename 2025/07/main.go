package main

import (
	_ "embed"
	"flag"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func part1(lines []string) (res int) {
	activeCols := make(map[int]bool)
	startFound := false

	for _, line := range lines {
		// Find start if not found
		if !startFound {
			idx := strings.Index(line, "S")
			if idx != -1 {
				activeCols[idx] = true
				startFound = true
			} else {
				continue
			}
		}

		nextCols := make(map[int]bool)
		for c := range activeCols {
			if c < 0 || c >= len(line) {
				continue
			}
			char := line[c]
			if char == '^' {
				res++
				nextCols[c-1] = true
				nextCols[c+1] = true
			} else {
				// '.' or 'S' or anything else passes through
				nextCols[c] = true
			}
		}
		activeCols = nextCols
		if len(activeCols) == 0 {
			break
		}
	}
	return
}

func part2(lines []string) (res int) {
	activeCols := make(map[int]int)
	startFound := false

	for _, line := range lines {
		// Find start if not found
		if !startFound {
			idx := strings.Index(line, "S")
			if idx != -1 {
				activeCols[idx] = 1
				startFound = true
			} else {
				continue
			}
		}

		nextCols := make(map[int]int)
		for c, count := range activeCols {
			// Check bounds
			if c < 0 || c >= len(line) {
				// Outside manifold, falls through
				nextCols[c] += count
				continue
			}

			char := line[c]
			if char == '^' {
				nextCols[c-1] += count
				nextCols[c+1] += count
			} else {
				// '.' or 'S' or anything else passes through
				nextCols[c] += count
			}
		}
		activeCols = nextCols
	}

	for _, count := range activeCols {
		res += count
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
