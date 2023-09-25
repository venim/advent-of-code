package main

import (
	_ "embed"
	"flag"
	"math"
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

type pos struct {
	X, Y, Z int
}

func parseFrequency(lines []string) []int {
	n := len(lines[0])
	ones := make([]int, n)
	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				ones[i]++
			}
		}
	}

	for i := range ones {
		ones[i] -= len(lines) / 2
	}
	return ones
}

func parseRates(frequency []int) (gamma int, epsilon int) {
	n := len(frequency)
	var g, e int
	for i := 0; i < n; i++ {
		b := math.Pow(2, float64(n-1-i))
		if frequency[i] > 0 {
			g += int(b)
		} else {
			e += int(b)
		}
	}
	return g, e
}

func part1(lines []string) (res int) {
	freq := parseFrequency(lines)
	gamma, epsilon := parseRates(freq)
	return int(gamma * epsilon)
}

func part2(lines []string) (res int) {
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
