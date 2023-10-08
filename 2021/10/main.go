package main

import (
	_ "embed"
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"golang.org/x/exp/slices"
)

var (
	//go:embed in.txt
	input string
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

func parseLine(s string) (remaining []rune, corrupt bool) {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'>': '<',
	}
	for _, s := range s {
		if _, ok := pairs[s]; ok {
			// closing bracket found
			if stack[len(stack)-1] != pairs[s] {
				// line is corrupt
				return []rune{s}, true
			}
			// close the pair
			stack = stack[:len(stack)-1]
		} else {
			// opening bracket found
			stack = append(stack, s)
		}
	}
	return stack, false
}

func part1(lines []string) (res int) {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, line := range lines {
		stack, corrupt := parseLine(line)
		if corrupt {
			res += points[stack[0]]
		}
	}
	return
}

func complete(stack []rune) string {
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
	rest := ""
	for _, s := range stack {
		rest = string(pairs[s]) + rest
	}
	return rest
}

func part2(lines []string) (res int) {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	scores := []int{}
	for _, line := range lines {
		stack, corrupt := parseLine(line)
		if !corrupt {
			s := complete(stack)
			score := 0
			for _, s := range s {
				score *= 5
				score += points[s]
			}
			scores = append(scores, score)
		}
	}
	slices.Sort(scores)
	return scores[len(scores)/2]
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
