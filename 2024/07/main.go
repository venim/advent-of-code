package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func eval(nums []int, ops []byte) int {
	n := nums[0]
	for i, op := range ops {
		switch op {
		case '+':
			n += nums[i+1]
		case '*':
			n *= nums[i+1]
		case '|':
			n = util.MustAtoi(fmt.Sprintf("%d%d", n, nums[i+1]))
		}
	}
	return n
}

func solve(want int, nums []int, eq []byte, ops []byte) bool {
	if len(eq) == len(nums)-1 {
		return eval(nums, eq) == want
	}
	for _, op := range ops {
		if solve(want, nums, append(eq, op), ops) {
			return true
		}
	}
	return false
}

func parse(line string) (int, []int) {
	parts := strings.Split(line, ":")
	want := util.MustAtoi(parts[0])
	nums := []int{}
	for _, n := range strings.Fields(parts[1]) {
		nums = append(nums, util.MustAtoi(n))
	}
	return want, nums
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		want, nums := parse(line)
		if solve(want, nums, []byte{}, []byte{'+', '*'}) {
			res += want
		}
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		want, nums := parse(line)
		if solve(want, nums, []byte{}, []byte{'+', '*', '|'}) {
			res += want
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
	lines := util.SplitLines(input)

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
