package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

type Point struct {
	R, C int
}

type Variant struct {
	Points []Point
	Masks  [4]uint64 // Increased to 4 for safety, though 3x3 is expected
	MaxR   int
	MaxC   int
	MinC   int
}

type Shape struct {
	ID       int
	Area     int
	Variants []Variant
}

type Region struct {
	W, H          int
	Counts        []int
	TotalArea     int
	TotalPresents int
}

func parse(lines []string) (shapes []Shape, regions []Region) {
	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			break
		}

		parts := strings.Split(line, ":")
		id, _ := strconv.Atoi(parts[0])
		i++
		var grid []string
		for i < len(lines) {
			l := strings.TrimRight(lines[i], " \r\n")
			if l == "" || (strings.Contains(l, "x") && strings.Contains(l, ":")) {
				break
			}
			grid = append(grid, l)
			i++
		}
		shapes = append(shapes, createShape(id, grid))
	}

	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		parts := strings.Split(line, ":")
		dims := strings.Split(parts[0], "x")
		w, _ := strconv.Atoi(dims[0])
		h, _ := strconv.Atoi(dims[1])
		countsStr := strings.Fields(parts[1])
		var counts []int
		totalArea := 0
		totalPresents := 0
		for idx, cs := range countsStr {
			c, _ := strconv.Atoi(cs)
			counts = append(counts, c)
			if idx < len(shapes) {
				totalArea += c * shapes[idx].Area
			}
			totalPresents += c
		}
		regions = append(regions, Region{W: w, H: h, Counts: counts, TotalArea: totalArea, TotalPresents: totalPresents})
		i++
	}
	return
}

func createShape(id int, grid []string) Shape {
	var points []Point
	for r, row := range grid {
		for c, char := range row {
			if char == '#' {
				points = append(points, Point{r, c})
			}
		}
	}

	area := len(points)
	variantsMap := make(map[string]Variant)

	curr := points
	for f := 0; f < 2; f++ {
		for r := 0; r < 4; r++ {
			v := normalize(curr)
			key := pointsKey(v.Points)
			variantsMap[key] = v
			// Rotate 90: (r, c) -> (c, -r)
			next := make([]Point, len(curr))
			for k, p := range curr {
				next[k] = Point{p.C, -p.R}
			}
			curr = next
		}
		// Flip: (r, c) -> (-r, c)
		next := make([]Point, len(points))
		for k, p := range points {
			next[k] = Point{-p.R, p.C}
		}
		curr = next
		points = curr
	}

	var variants []Variant
	for _, v := range variantsMap {
		variants = append(variants, v)
	}
	return Shape{ID: id, Area: area, Variants: variants}
}

func normalize(points []Point) Variant {
	minR, minC := 100, 100
	for _, p := range points {
		if p.R < minR { minR = p.R }
		if p.C < minC { minC = p.C }
	}
	res := make([]Point, len(points))
	maxR, maxC, realMinC := 0, 0, 100
	for i, p := range points {
		res[i] = Point{p.R - minR, p.C - minC}
		if res[i].R > maxR { maxR = res[i].R }
		if res[i].C > maxC { maxC = res[i].C }
		if res[i].C < realMinC { realMinC = res[i].C }
	}
	// Sort points lexicographically
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i].R > res[j].R || (res[i].R == res[j].R && res[i].C > res[j].C) {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	var masks [4]uint64
	for _, p := range res {
		masks[p.R] |= (1 << uint(p.C))
	}
	return Variant{Points: res, Masks: masks, MaxR: maxR, MaxC: maxC, MinC: realMinC}
}

func pointsKey(pts []Point) string {
	var s string
	for _, p := range pts {
		s += fmt.Sprintf("%d,%d;", p.R, p.C)
	}
	return s
}

func solve(reg Region, shapes []Shape) bool {
	if reg.TotalArea > reg.W*reg.H {
		return false
	}
	grid := [50]uint64{}
	for r := 0; r < 50; r++ {
		if r < reg.H {
			grid[r] = (^uint64(0)) << uint(reg.W)
		} else {
			grid[r] = ^uint64(0)
		}
	}

	return dfs(0, 0, &grid, reg.W, reg.H, reg.Counts, reg.TotalArea, reg.W*reg.H, shapes)
}

func dfs(r, c int, grid *[50]uint64, W, H int, counts []int, remainingArea int, availableCells int, shapes []Shape) bool {
	if remainingArea == 0 {
		return true
	}
	if availableCells < remainingArea {
		return false
	}

	for r < H {
		if (grid[r] & (1 << uint(c))) == 0 {
			break
		}
		c++
		if c == W {
			c = 0
			r++
		}
	}
	if r == H {
		return remainingArea == 0
	}

	for i := range shapes {
		if counts[i] > 0 {
			for _, v := range shapes[i].Variants {
				for _, p := range v.Points {
					or, oc := r-p.R, c-p.C
					if or < 0 || or+v.MaxR >= H || oc < 0 || oc+v.MaxC >= W {
						continue
					}

					overlap := false
					for dr := 0; dr <= v.MaxR; dr++ {
						if (grid[or+dr] & (v.Masks[dr] << uint(oc))) != 0 {
							overlap = true
							break
						}
					}
					if overlap {
						continue
					}

					for dr := 0; dr <= v.MaxR; dr++ {
						grid[or+dr] |= (v.Masks[dr] << uint(oc))
					}
					counts[i]--
					if dfs(r, c, grid, W, H, counts, remainingArea-shapes[i].Area, availableCells-1, shapes) {
						return true
					}
					counts[i]++
					for dr := 0; dr <= v.MaxR; dr++ {
						grid[or+dr] &= ^(v.Masks[dr] << uint(oc))
					}
				}
			}
		}
	}

	if availableCells > remainingArea {
		grid[r] |= (1 << uint(c))
		if dfs(r, c, grid, W, H, counts, remainingArea, availableCells-1, shapes) {
			return true
		}
		grid[r] &= ^(1 << uint(c))
	}

	return false
}

func part1(lines []string) (res int) {
	shapes, regions := parse(lines)
	for _, reg := range regions {
		if solve(reg, shapes) {
			res++
		}
	}
	return
}

func part2(lines []string) (res int) {
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
