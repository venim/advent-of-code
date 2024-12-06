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

func parse(input string) (map[int][]int, []map[int]int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	rules := map[int][]int{}
	updates := []map[int]int{}
	parts := strings.Split(input, "\n\n")
	for _, rule := range strings.Split(parts[0], "\n") {
		parts := strings.Split(rule, "|")
		x := util.MustAtoi(parts[0])
		y := util.MustAtoi(parts[1])
		if _, ok := rules[y]; !ok {
			rules[y] = []int{x}
		} else {
			rules[y] = append(rules[y], x)
		}
	}
	for _, line := range strings.Split(parts[1], "\n") {
		update := map[int]int{}
		for i, n := range strings.Split(line, ",") {
			update[util.MustAtoi(n)] = i
		}
		updates = append(updates, update)
	}
	return rules, updates
}

func validOrder(rules map[int][]int, update map[int]int) bool {
	for n, i := range update {
		afters := rules[n]
		for _, n := range afters {
			n, ok := update[n]
			if ok && n > i {
				return false
			}
		}
	}
	return true
}

func middlePage(update map[int]int) int {
	u := make([]int, len(update))
	for n, i := range update {
		u[i] = n
	}
	return u[len(update)/2]
}

func part1(input string) (res int) {
	rules, updates := parse(input)
	for _, update := range updates {
		if !validOrder(rules, update) {
			continue
		}
		res += middlePage(update)
	}
	return
}

func findFirst(rules map[int][]int, update map[int]int) int {
loop:
	for n := range update {
		for _, after := range rules[n] {
			if _, ok := update[after]; ok {
				continue loop
			}
		}
		delete(update, n)
		return n
	}
	return 0
}

func reorder(afters map[int][]int, update map[int]int) map[int]int {
	u := map[int]int{}
	i := 0

	for len(update) > 0 {
		n := findFirst(afters, update)
		u[n] = i
		i++
	}
	return u
}

func part2(input string) (res int) {
	afters, updates := parse(input)
	for _, update := range updates {
		if !validOrder(afters, update) {
			update = reorder(afters, update)
			res += middlePage(update)
		}
	}
	return
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

	t = time.Now()
	res = part1(input)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(input)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
