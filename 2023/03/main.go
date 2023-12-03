package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func part1(lines []string) (res int) {
	symbols := map[util.Pos]bool{}
	symbolsRe := regexp.MustCompile(`[^\d\.]`)
	for i, line := range lines {
		found := symbolsRe.FindAllStringIndex(line, -1)
		for _, f := range found {
			symbols[util.Pos{X: f[0], Y: i}] = true
		}
	}

	numRe := regexp.MustCompile(`(\d+)`)
	for i, line := range lines {
		found := numRe.FindAllStringIndex(line, -1)
	loop:
		for _, f := range found {
			for y := i - 1; y <= i+1; y++ {
				for x := f[0] - 1; x <= f[1]; x++ {
					if y < 0 || y >= len(lines) || x < 0 || x >= len(line) ||
						(y == i && (x >= f[0] && x < f[1])) {
						continue
					}
					if symbols[util.Pos{X: x, Y: y}] {
						res += util.MustAtoi(line[f[0]:f[1]])
						continue loop
					}
				}
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	gearsRe := regexp.MustCompile(`(\*)`)
	numRe := regexp.MustCompile(`(\d+)`)
	for i, line := range lines {
		gears := gearsRe.FindAllStringIndex(line, -1)
		if len(gears) == 0 {
			continue
		}
		for _, gear := range gears {
			parts := []int{}
			for y := i - 1; y <= i+1; y++ {
				if y < 0 || y >= len(lines) {
					continue
				}
				nums := numRe.FindAllStringIndex(lines[y], -1)
			numLoop:
				for _, num := range nums {
					for _, x := range []int{num[0], num[1] - 1} {
						if x >= gear[0]-1 && x <= gear[0]+1 {
							parts = append(parts, util.MustAtoi(lines[y][num[0]:num[1]]))
							continue numLoop
						}
					}
				}
			}
			if len(parts) == 2 {
				product := 1
				for _, p := range parts {
					product *= p
				}
				res += product
			}
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
