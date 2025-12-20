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

func parseGraph(lines []string) map[string][]string {
	adj := make(map[string][]string)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		src := parts[0]
		dsts := strings.Fields(parts[1])
		adj[src] = dsts
	}
	return adj
}

func countPaths(u, target string, adj map[string][]string, memo map[string]int) int {
	if u == target {
		return 1
	}
	if v, ok := memo[u]; ok {
		return v
	}
	sum := 0
	if neighbors, ok := adj[u]; ok {
		for _, v := range neighbors {
			sum += countPaths(v, target, adj, memo)
		}
	}
	memo[u] = sum
	return sum
}

func part1(lines []string) (res int) {
	adj := parseGraph(lines)
	return countPaths("you", "out", adj, make(map[string]int))
}

func part2(lines []string) (res int) {
	adj := parseGraph(lines)

	// Path 1: svr -> dac -> fft -> out
	p1_1 := countPaths("svr", "dac", adj, make(map[string]int))
	p1_2 := countPaths("dac", "fft", adj, make(map[string]int))
	p1_3 := countPaths("fft", "out", adj, make(map[string]int))
	path1 := p1_1 * p1_2 * p1_3

	// Path 2: svr -> fft -> dac -> out
	p2_1 := countPaths("svr", "fft", adj, make(map[string]int))
	p2_2 := countPaths("fft", "dac", adj, make(map[string]int))
	p2_3 := countPaths("dac", "out", adj, make(map[string]int))
	path2 := p2_1 * p2_2 * p2_3

	return path1 + path2
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
