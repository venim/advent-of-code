package main

import (
	_ "embed"
	"flag"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

var numRe = regexp.MustCompile(`\d+`)

func parseNumbers(line string) (nums []int) {
	matches := numRe.FindAllString(line, -1)
	for _, n := range matches {
		nums = append(nums, util.MustAtoi(n))
	}
	return
}

var operRe = regexp.MustCompile(`[\+\*]`)

func parseOperands(line string) (opers []string) {
	return operRe.FindAllString(line, -1)
}

func part1(lines []string) (res int) {
	var numbers [][]int
	var operands []string
	for i, line := range lines {
		if i == len(lines)-1 {
			operands = parseOperands(line)
			continue
		}
		numbers = append(numbers, parseNumbers(line))
	}

	for i := range len(operands) {
		var n int
		for j := range len(numbers) {
			if j == 0 {
				n = numbers[j][i]
				continue
			}
			switch operands[i] {
			case "*":
				n *= numbers[j][i]
			case "+":
				n += numbers[j][i]
			}
		}
		res += n
	}
	return
}

func parseColumns(in []string) (out []int) {
	var N int
	for _, s := range in {
		if len(s) > N {
			N = len(s)
		}
	}
	for i := range N {
		var n []byte
		for _, s := range in {
			if i < len(s) && s[i] != ' ' {
				n = append(n, s[i])
			}
		}
		out = append(out, util.MustAtoi(string(n)))
	}
	slices.Reverse(out)
	return
}

var part2Re = regexp.MustCompile(`[\+\*]\s+`)

func part2(lines []string) (res int) {
	var numbers [][]string
	var operands []string

	for range len(lines) - 1 {
		numbers = append(numbers, []string{})
	}

	for lines[0] != "" {
		// parse each column one by one
		line := lines[len(lines)-1]
		loc := part2Re.FindStringIndex(line)
		operands = append(operands, line[:1])

		for i, line := range lines {
			lines[i] = line[loc[1]:]
			if i == len(lines)-1 {
				continue
			}
			n := line[:loc[1]]
			if loc[1] != len(line) {
				n = n[:len(n)-1]
			}
			numbers[i] = append(numbers[i], n)
		}
	}

	for i := range len(numbers[0]) {
		var ns []string
		for _, n := range numbers {
			ns = append(ns, n[i])
		}
		out := parseColumns(ns)
		n := out[0]
		for _, nn := range out[1:] {
			switch operands[i] {
			case "*":
				n *= nn
			case "+":
				n += nn
			}
		}
		res += n
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
