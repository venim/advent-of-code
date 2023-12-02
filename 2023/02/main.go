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

func analyzeGame(line string, cubeCallbackFn func(color string, n int) (ok bool)) int {
	parts := strings.Split(line, ": ")
	sets := strings.Split(parts[1], "; ")
	for _, set := range sets {
		for _, cube := range strings.Split(set, ", ") {
			cube := strings.Split(cube, " ")
			if ok := cubeCallbackFn(cube[1], util.MustAtoi(cube[0])); !ok {
				return 0
			}
		}
	}
	id, _ := strings.CutPrefix(parts[0], "Game ")
	return util.MustAtoi(id)
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		res += analyzeGame(line, func(c string, n int) bool {
			switch c {
			case "red":
				if n > 12 {
					return false
				}
			case "blue":
				if n > 14 {
					return false
				}
			case "green":
				if n > 13 {
					return false
				}
			}
			return true
		})
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		minCubes := map[string]int{}
		analyzeGame(line, func(c string, n int) bool {
			if minCubes[c] < n {
				minCubes[c] = n
			}
			return true
		})

		power := 1
		for _, n := range minCubes {
			power *= n
		}
		res += power
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
