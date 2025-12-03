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

func findJoltage(bank string, n int) int {
	if len(bank) <= n { // If the bank is already 12 digits or less, we keep all of them.
		return util.MustAtoi(bank)
	}

	nToDrop := len(bank) - n
	stack := []rune{}

	for _, digit := range bank {
		for len(stack) > 0 && // if there's any digit left
			digit > stack[len(stack)-1] && // if the current digit is greater than the current end
			nToDrop > 0 { // if we still have to drop a battery
			stack = stack[:len(stack)-1] // Pop
			nToDrop--
		}
		stack = append(stack, digit) // Push
	}

	// If nToDrop is still > 0, it means the remaining digits are in non-increasing order.
	// We need to remove from the end to keep the most significant digits.
	stack = stack[:len(stack)-nToDrop]

	// Ensure the result has exactly 12 digits.
	if len(stack) > n {
		stack = stack[:n]
	}

	return util.MustAtoi(string(stack))
}

func part1(banks []string) (res int) {
	for _, b := range banks {
		res += findJoltage(b, 2)
	}
	return
}

func part2(banks []string) (res int) {
	for _, b := range banks {
		res += findJoltage(b, 12)
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
