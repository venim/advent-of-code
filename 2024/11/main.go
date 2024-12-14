package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
	cache = map[string]int{}
)

func blink(stone int, remaining int) int {
	cacheKey := fmt.Sprintf("%d-%d", stone, remaining)
	if _, ok := cache[cacheKey]; ok {
	} else if remaining == 0 {
		cache[cacheKey] = 1
	} else if stone == 0 {
		cache[cacheKey] = blink(1, remaining-1)
	} else if n := int(math.Floor(math.Log10(float64(stone))) + 1); n%2 == 0 {
		d := int(math.Pow10(n / 2))
		cache[cacheKey] = blink(stone/d, remaining-1) + blink(stone%d, remaining-1)
	} else {
		cache[cacheKey] = blink(stone*2024, remaining-1)
	}

	return cache[cacheKey]
}

func part1(lines []string) (res int) {
	stones := strings.Fields(lines[0])
	for _, stone := range stones {
		res += blink(util.MustAtoi(stone), 25)
	}
	return
}

func part2(lines []string) (res int) {
	stones := strings.Fields(lines[0])
	for _, stone := range stones {
		res += blink(util.MustAtoi(stone), 75)
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
