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

// a parseFn parses the space-separated fields of an instruction
// and returns the number of meters to travel and the direction vector
type parseFn func(fields []string) (n int, dir complex128)

// findArea finds the area of a polygon using the Shoelace formula
// https://en.wikipedia.org/wiki/Shoelace_formula
// the vertices in the slice must be contiguous along the polygon
func findArea(vertices []complex128) (area float64) {
	n := len(vertices)
	for i := 0; i < n; i++ {
		v := vertices[i]
		vv := vertices[(i+1)%n]
		// trapezoid formula
		area += real(v-vv) * imag(v+vv)
	}
	return math.Abs(area) / 2
}

func solve(instructions []string, parse parseFn) int {
	var (
		v         complex128
		perimeter int
		vertices  = make([]complex128, 0, len(instructions))
	)

	// follow instructions to find perimeter and all vertices
	for _, line := range instructions {
		n, d := parse(strings.Fields(line))
		perimeter += n
		v += d * complex(float64(n), 0)
		vertices = append(vertices, v)
	}

	area := findArea(vertices)

	// use Pick's theorem to find # of inner interior
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	// i = A - b/2 + 1
	interior := int(area - float64(perimeter)*0.5 + 1)
	return perimeter + interior
}

func part1(lines []string) (res int) {
	decoder := map[string]complex128{"R": 1, "L": -1, "U": -1i, "D": 1i}

	return solve(lines, func(fields []string) (n int, dir complex128) {
		n = util.MustAtoi(fields[1])

		dir = decoder[fields[0]]
		return
	})
}

func part2(lines []string) (res int) {
	decoder := map[byte]complex128{'0': 1, '1': 1i, '2': -1, '3': -1i}

	return solve(lines, func(fields []string) (n int, dir complex128) {
		// strip ()#
		hex := strings.NewReplacer("(", "", ")", "", "#", "").Replace(fields[2])

		// convert the first 5 digits
		n64, _ := strconv.ParseInt(hex[:5], 16, 64)
		n = int(n64)

		// decode with the last digit
		dir = decoder[hex[5]]
		return
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
