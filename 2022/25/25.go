package main

import (
	_ "embed"
	"flag"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"golang.org/x/exp/slices"
)

var (
	//go:embed in.txt
	input  string
	digits = []string{"=", "-", "0", "1", "2"}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func stoi(s string) (i int) {
	for s != "" {
		n := (slices.Index(digits, s[:1]) - 2)
		pow := float64(len(s) - 1)
		i += n * int(math.Pow(5, pow))
		s = s[1:]
	}
	return
}

func itos(i int) (s string) {
	for i > 0 {
		b := (i + 2) % 5
		s = digits[b] + s
		i = (i + 2) / 5
	}
	return
}

func part1(lines []string) (res string) {
	sum := 0
	for _, l := range lines {
		sum += stoi(l)
	}
	return itos(sum)
}

func part2(lines []string) (res string) {
	return
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t   time.Time
		res string
	)
	flag.Parse()
	lines := strings.Split(input, "\n")

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
