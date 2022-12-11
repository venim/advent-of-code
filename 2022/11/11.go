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

	ops map[string]func(a, b int) int = map[string]func(a int, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type monkey struct {
	Items  []int
	Op     func(int) int
	Tester int
	Dest   map[bool]int
}

func makeMonkeys(lines []string) (int, []*monkey) {
	monkeys := []*monkey{}
	cm := 1 // common multiple used to control worry level
	for i := 0; i < len(lines); i++ {
		monkey := &monkey{}

		// skip id parsing since input is in order
		// id := mustAtoi(lines[i][7:8])
		i++

		monkey.Items = parseItems(lines[i])
		i++

		monkey.Op = makeOp(lines[i])
		i++

		monkey.Tester = makeTester(lines[i])
		cm *= monkey.Tester
		i++

		t := parseDest(lines[i])
		i++

		f := parseDest(lines[i])
		i++

		monkey.Dest = map[bool]int{true: t, false: f}
		monkeys = append(monkeys, monkey)
	}
	return cm, monkeys
}

func parseItems(line string) []int {
	items := strings.Split(strings.Split(line, ": ")[1], ", ")
	res := make([]int, len(items))
	for i, item := range items {
		res[i] = mustAtoi(item)
	}
	return res
}

func makeOp(line string) func(int) int {
	parts := strings.Split(strings.Split(line, ": ")[1], " ")
	a, aErr := strconv.Atoi(parts[2])
	b, bErr := strconv.Atoi(parts[4])
	op := ops[parts[3]]
	return func(old int) int {
		if aErr != nil {
			a = old
		}
		if bErr != nil {
			b = old
		}
		return op(a, b)
	}
}

func makeTester(line string) int {
	return mustAtoi(strings.Split(line, "divisible by ")[1])
}

func parseDest(line string) int {
	return mustAtoi(strings.Split(line, "throw to monkey ")[1])
}

func solve(lines []string, rounds int, part int) (res int) {
	cm, monkeys := makeMonkeys(lines)
	inspections := make([]int, len(monkeys))
	for round := 0; round < rounds; round++ {
		for n, m := range monkeys {
			for _, i := range m.Items {
				inspections[n]++
				i = m.Op(i)
				switch part {
				case 1:
					i /= 3
				case 2:
					i %= cm
				}
				dest := m.Dest[i%m.Tester == 0]
				monkeys[dest].Items = append(monkeys[dest].Items, i)
			}
			m.Items = []int{}
		}
	}
	sort.Ints(inspections)
	res = inspections[len(inspections)-1] * inspections[len(inspections)-2]

	return
}

func main() {
	var (
		t time.Time
	)
	lines := strings.Split(input, "\n")

	t = time.Now()
	part1 := solve(lines, 20, 1)
	fmt.Printf("[Part 1] = %v\n", part1)
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	part2 := solve(lines, 10000, 2)
	fmt.Printf("[Part 2] = \n%v\n", part2)
	fmt.Printf("took %s\n\n", time.Since(t))
}
