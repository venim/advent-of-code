package main

import (
	_ "embed"
	"flag"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func solve(lines []string, fn func(line string) (n int, dir complex128)) int {
	outer := 0
	vertices := make([]complex128, 0, len(lines))
	var v complex128
	for _, line := range lines {
		n, d := fn(line)
		outer += n
		v += d * complex(float64(n), 0)
		vertices = append(vertices, v)
	}

	var area float64

	n := len(lines)
	j := n - 1
	for i := 0; i < n; i++ {
		v := vertices[i]
		vv := vertices[j]
		area += real(v)*imag(vv) - real(vv)*imag(v)
		j = i
	}
	area = math.Abs(area) / 2

	// Pick's theorem to find # of inner points
	// A - b/2 + 1 = i
	inner := int(area - float64(outer)*0.5 + 1)
	return outer + inner
}

func part1(lines []string) (res int) {
	return solve(lines, func(s string) (int, complex128) {
		parts := strings.Fields(s)
		n := util.MustAtoi(parts[1])
		var d complex128
		switch parts[0] {
		case "R":
			d = 1
		case "L":
			d = -1
		case "U":
			d = -1i
		case "D":
			d = 1i
		}
		return n, d
	})
}

func part2(lines []string) (res int) {
	return solve(lines, func(line string) (int, complex128) {
		parts := strings.Fields(line)
		hex := parts[2][2 : len(parts[2])-1]
		n, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 64)
		var d complex128
		switch hex[len(hex)-1] {
		case '0':
			d = 1
		case '2':
			d = -1
		case '3':
			d = -1i
		case '1':
			d = 1i
		}
		return int(n), d
	})
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
