package main

import (
	_ "embed"
	"flag"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input  string
	dirs   = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

type template struct {
	items map[string]int
	pairs map[string]int
	rules map[string]string
}

func (t *template) run() {
	newPairs := map[string]int{}
	for pair, n := range t.pairs {
		if rule, ok := t.rules[pair]; ok {
			newPairs[string(pair[0])+rule] += n
			newPairs[rule+string(pair[1])] += n
			t.items[rule] += n
		} else {
			// this apparently never happens,
			// but we should add pairs without a match to the next evolution
			newPairs[pair] += n
		}
	}
	t.pairs = newPairs
}

func (t *template) score() int {
	maxN := -1
	minN := math.MaxInt

	for _, n := range t.items {
		maxN = max(maxN, n)
		minN = min(minN, n)
	}

	return maxN - minN
}

func parse(lines []string) *template {
	pairs := map[string]int{}
	items := map[string]int{}
	rules := map[string]string{}

	start := lines[0]

	for i := 0; i < len(start)-1; i++ {
		// count all pairs and the first character
		pairs[start[i:i+2]]++
		items[start[i:i+1]]++
	}
	// count the last character
	items[start[len(start)-1:]]++

	lines = lines[2:]
	for _, l := range lines {
		pair, insert, _ := strings.Cut(l, " -> ")
		rules[pair] = insert
	}
	return &template{
		pairs: pairs,
		items: items,
		rules: rules,
	}
}

func part1(lines []string) (res int) {
	template := parse(lines)
	for i := 0; i < 10; i++ {
		template.run()
	}
	return template.score()
}

func part2(lines []string) (res int) {
	template := parse(lines)
	for i := 0; i < 40; i++ {
		template.run()
	}
	return template.score()
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
