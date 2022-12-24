package main

import (
	_ "embed"
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func op(op string, monkeys map[string]*monkey, m ...string) *int {
	var res int
	switch op {
	case "+":
		res = *monkeys[m[0]].N + *monkeys[m[1]].N
	case "-":
		res = *monkeys[m[0]].N - *monkeys[m[1]].N
	case "*":
		res = *monkeys[m[0]].N * *monkeys[m[1]].N
	case "/":
		res = *monkeys[m[0]].N / *monkeys[m[1]].N
	default:
		return nil
	}
	return &res
}

func inv(op string, c, b, i int) int {
	switch op {
	case "+":
		return c - b
	case "-":
		if i == 0 {
			return b - c
		}
		return c + b
	case "*":
		return c / b
	case "/":
		return c * b
	}
	return -1
}

type monkey struct {
	ID string
	// use pointer to int to enable N==nil check to determine if the monkey knows
	// what value to shout, or if it depends on others
	N       *int
	Monkeys []string
	Op      string
	X       bool
}

func (m *monkey) calc(monkeys map[string]*monkey) {
	if m.X {
		return
	}
	if m.Monkeys != nil {
		for _, c := range m.Monkeys {
			monkeys[c].calc(monkeys)
		}
		for _, c := range m.Monkeys {
			if monkeys[c].N == nil {
				return
			}
		}
		m.N = op(m.Op, monkeys, m.Monkeys...)
	}
}

func (m *monkey) findX(monkeys map[string]*monkey, want int) int {
	if m.X {
		return want
	}
	for i, c := range m.Monkeys {
		mc := monkeys[c]
		if monkeys[c].N != nil {
			want = inv(m.Op, want, *mc.N, i)
			return monkeys[m.Monkeys[(i+1)%2]].findX(monkeys, want)
		}
	}
	return -1
}

func parseMonkey(line string) *monkey {
	parts := strings.Split(line, ": ")
	m := &monkey{
		ID: parts[0],
	}
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		parts = strings.Split(parts[1], " ")
		m.Monkeys = []string{parts[0], parts[2]}
		m.Op = parts[1]
	} else {
		m.N = &n
	}
	return m
}

func part1(lines []string) (res int) {
	monkeys := make(map[string]*monkey, len(lines))
	for _, line := range lines {
		m := parseMonkey(line)
		monkeys[m.ID] = m
	}
	monkeys["root"].calc(monkeys)
	return *monkeys["root"].N
}

func part2(lines []string) (res int) {
	monkeys := make(map[string]*monkey, len(lines))
	for _, line := range lines {
		m := parseMonkey(line)
		monkeys[m.ID] = m
	}
	monkeys["humn"].X = true
	monkeys["humn"].N = nil
	m := monkeys["root"]
	m.calc(monkeys)
	for i, c := range m.Monkeys {
		if monkeys[c].N != nil {
			return monkeys[m.Monkeys[(i+1)%2]].findX(monkeys, *monkeys[c].N)
		}
	}
	return -1
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
