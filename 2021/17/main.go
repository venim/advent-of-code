package main

import (
	_ "embed"
	"flag"
	"math"
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

type trajectory struct {
	maxY int
	pos  util.Pos
	v    util.Pos
}

func (t *trajectory) step() {
	t.pos.X += t.v.X
	t.pos.Y += t.v.Y
	t.maxY = max(t.maxY, t.pos.Y)
	if t.v.X > 0 {
		t.v.X--
	} else if t.v.X < 0 {
		t.v.X++
	}
	t.v.Y--
}

func (t *trajectory) willLand(target target) bool {
	for {
		t.step()

		if t.pos.X >= target.min.X && t.pos.X <= target.max.X &&
			t.pos.Y >= target.min.Y && t.pos.Y <= target.max.Y {
			return true
		}

		if t.pos.Y < target.min.Y {
			return false
		}
		if t.pos.X > target.max.X {
			return false
		}
	}
}

type target struct {
	min util.Pos
	max util.Pos
}

func run(targetLine string, onLand func(t *trajectory)) {
	re := regexp.MustCompile(`x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	match := re.FindStringSubmatch(targetLine)
	minPos := util.Pos{X: util.MustAtoi(match[1]), Y: util.MustAtoi(match[3])}
	maxPos := util.Pos{X: util.MustAtoi(match[2]), Y: util.MustAtoi(match[4])}

	target := target{minPos, maxPos}
	for x := 0; x < target.max.X+1; x++ {
		for y := target.min.Y; y < -1*target.min.Y+1; y++ {
			t := &trajectory{maxY: -1 * math.MaxInt, v: util.Pos{X: x, Y: y}}
			if t.willLand(target) {
				onLand(t)
			}
		}
	}
}

func part1(lines []string) (res int) {
	maxY := -1 * math.MaxInt
	run(lines[0], func(t *trajectory) {
		maxY = max(maxY, t.maxY)
	})

	return maxY
}

func part2(lines []string) (res int) {
	run(lines[0], func(t *trajectory) {
		res++
	})

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
