package main

import (
	_ "embed"
	"flag"
	"math"
	"slices"
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
		ones[i] -= int(math.Ceil(float64(len(lines)) / 2))
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

func findRating(lines []string, bitCriteria func(int) byte) int {
	numbers := slices.Clone(lines)
	for i := 0; i < len(numbers[0]); i++ {
		freq := parseFrequency(numbers)
		remaining := make([]string, 0, len(lines))
		want := bitCriteria(freq[i])
		for _, number := range numbers {
			if number[i] == want {
				remaining = append(remaining, number)
			}
		}
		if len(remaining) == 1 {
			number := remaining[0]
			rating, err := strconv.ParseInt(number, 2, 0)
			if err != nil {
				return 0
			}
			return int(rating)
		}
		numbers = remaining
	}
	return 0
}

func findOxygenGeneratorRating(lines []string) int {
	return findRating(lines, func(i int) byte {
		if i >= 0 {
			return '1'
		}
		return '0'
	})
}

func findCO2ScrubberRating(lines []string) int {
	return findRating(lines, func(i int) byte {
		if i < 0 {
			return '1'
		}
		return '0'
	})
}

func part2(lines []string) (res int) {
	ogr := findOxygenGeneratorRating(lines)
	co2sr := findCO2ScrubberRating(lines)

	return ogr * co2sr
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
