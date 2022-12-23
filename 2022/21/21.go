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

type monkey struct {
	ID      string
	N       int
	Monkeys []string
	Op      string
}

func (m *monkey) calc(monkeys map[string]*monkey) {
	if m.Monkeys != nil {
		m1 := monkeys[m.Monkeys[0]]
		m2 := monkeys[m.Monkeys[1]]
		m1.calc(monkeys)
		m2.calc(monkeys)
		switch m.Op {
		case "+":
			m.N = m1.N + m2.N
		case "-":
			m.N = m1.N - m2.N
		case "*":
			m.N = m1.N * m2.N
		case "/":
			m.N = m1.N / m2.N
		}
	}
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
		m.N = n
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
	return monkeys["root"].N
}

func part2(lines []string) (res int) {
	monkeys := make(map[string]*monkey, len(lines))
	for _, line := range lines {
		m := parseMonkey(line)
		monkeys[m.ID] = m
	}

	// for i := 1; i < math.MaxInt; i++ {
	// 	monkeys["humn"].N = i
	// 	m := monkeys["root"]
	// 	m.calc(monkeys)
	// 	if monkeys[m.Monkeys[0]].N == monkeys[m.Monkeys[1]].N {
	// 		return i
	// 	}
	// }
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
