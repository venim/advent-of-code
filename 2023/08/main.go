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

type Node struct {
	left  string
	right string
}

type Nodes map[string]*Node

func (nodes Nodes) new(id string) *Node {
	node, ok := nodes[id]
	if ok {
		return node
	}
	node = &Node{}
	nodes[id] = node
	return node
}

func (nodes Nodes) run(start string, end func(string) bool, instructions string) int {
	cur := start
	steps := 0
	i := 0
	for !end(cur) {
		steps++
		switch instructions[i] {
		case 'R':
			cur = nodes[cur].right
		case 'L':
			cur = nodes[cur].left
		}
		i++
		i %= len(instructions)
	}
	return steps
}

func parse(lines []string) Nodes {
	nodes := Nodes{}
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		n := parts[0]
		parts = strings.Split(parts[1], ", ")
		l := parts[0][1:]
		r := parts[1][:3]

		node := nodes.new(n)
		node.left = l
		node.right = r
	}
	return nodes
}

func part1(lines []string) (res int) {
	instructions := lines[0]
	nodes := parse(lines[2:])

	return nodes.run(
		"AAA",
		func(s string) bool { return s == "ZZZ" },
		instructions,
	)
}

func part2(lines []string) (res int) {
	instructions := lines[0]
	nodes := parse(lines[2:])

	starts := []string{}
	for i := range nodes {
		if strings.HasSuffix(i, "A") {
			starts = append(starts, i)
		}
	}

	cycles := []int{}
	end := func(s string) bool { return strings.HasSuffix(s, "Z") }
	for _, start := range starts {
		steps := nodes.run(start, end, instructions)
		cycles = append(cycles, steps)
	}

	return lcm(cycles...)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers ...int) int {
	a := integers[0]
	b := integers[1]
	integers = integers[2:]
	res := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		res = lcm(res, integers[i])
	}
	return res
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
