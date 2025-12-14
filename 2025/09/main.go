package main

import (
	_ "embed"
	"flag"
	"math/bits"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

// parseInput converts the input lines into a slice of positions and finds the maximum Y coordinate.
func parseInput(lines []string) (tiles []util.Pos, maxY int) {
	for _, l := range lines {
		p := strings.Split(l, ",")
		y := util.MustAtoi(p[1])
		if y > maxY {
			maxY = y
		}
		tiles = append(tiles, util.Pos{X: util.MustAtoi(p[0]), Y: y})
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func calcArea(a, b util.Pos) int {
	return (abs(a.X-b.X) + 1) * (abs(a.Y-b.Y) + 1)
}

func part1(lines []string) (res int) {
	tiles, _ := parseInput(lines)
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			if a := calcArea(tiles[i], tiles[j]); a > res {
				res = a
			}
		}
	}
	return
}

// PolygonChecker efficiently checks if a vertical range is valid within the polygon.
type PolygonChecker struct {
	stMin *SparseTable // Range Max Query on Left Walls
	stMax *SparseTable // Range Min Query on Right Walls
}

func NewPolygonChecker(tiles []util.Pos, maxY int) *PolygonChecker {
	// 1. Initialize Scanlines
	const inf = int(^uint(0) >> 1)
	const negInf = -inf - 1

	minXs := make([]int, maxY+1)
	maxXs := make([]int, maxY+1)
	for i := range minXs {
		minXs[i] = inf
		maxXs[i] = negInf
	}

	// 2. Fill Scanlines (Trace edges)
	n := len(tiles)
	for i := range n {
		p1, p2 := tiles[i], tiles[(i+1)%n]

		if p1.X == p2.X { // Vertical Edge
			yStart, yEnd := p1.Y, p2.Y
			if yStart > yEnd {
				yStart, yEnd = yEnd, yStart
			}
			for y := yStart; y <= yEnd; y++ {
				if p1.X < minXs[y] {
					minXs[y] = p1.X
				}
				if p1.X > maxXs[y] {
					maxXs[y] = p1.X
				}
			}
		} else { // Horizontal Edge
			// For horizontal edges, the X range is valid at this specific Y.
			// However, the "walls" logic is primarily defined by vertical boundaries.
			// A horizontal edge connects two vertical walls.
			// We strictly update the min/max for the single row Y.
			xStart, xEnd := p1.X, p2.X
			if xStart > xEnd {
				xStart, xEnd = xEnd, xStart
			}
			y := p1.Y
			if xStart < minXs[y] {
				minXs[y] = xStart
			}
			if xEnd > maxXs[y] {
				maxXs[y] = xEnd
			}
		}
	}

	// 3. Build Sparse Tables
	// We want the Rightmost Left-Wall (Max of minXs)
	// We want the Leftmost Right-Wall (Min of maxXs)
	return &PolygonChecker{
		stMin: NewSparseTable(minXs, func(a, b int) int { return max(a, b) }),
		stMax: NewSparseTable(maxXs, func(a, b int) int { return min(a, b) }),
	}
}

// IsValidRect checks if the rectangle defined by x1, x2 and y1, y2 is strictly inside the polygon.
// It assumes x1 <= x2 and y1 <= y2.
func (pc *PolygonChecker) IsValidRect(x1, x2, y1, y2 int) bool {
	// The polygon narrows to [limitL, limitR] over the range [y1, y2].
	// Our rectangle [x1, x2] must fit inside this narrowest point.
	limitL := pc.stMin.Query(y1, y2)
	limitR := pc.stMax.Query(y1, y2)
	return limitL <= x1 && x2 <= limitR
}

func part2(lines []string) (res int) {
	tiles, maxY := parseInput(lines)
	pc := NewPolygonChecker(tiles, maxY)

	for i := range tiles {
		p1 := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			p2 := tiles[j]

			// Pre-calculate dimensions
			w := p1.X - p2.X
			if w < 0 {
				w = -w
			}
			w++

			h := p1.Y - p2.Y
			if h < 0 {
				h = -h
			}
			h++

			area := w * h
			if area <= res {
				continue
			}

			// Sort coords for the check
			x1, x2 := p1.X, p2.X
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			y1, y2 := p1.Y, p2.Y
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			if pc.IsValidRect(x1, x2, y1, y2) {
				res = area
			}
		}
	}
	return
}

// --- Sparse Table Implementation ---

type SparseTable struct {
	data [][]int
	op   func(int, int) int
}

func NewSparseTable(arr []int, op func(int, int) int) *SparseTable {
	n := len(arr)
	k := bits.Len(uint(n))
	data := make([][]int, k)
	data[0] = make([]int, n)
	copy(data[0], arr)

	for i := 1; i < k; i++ {
		data[i] = make([]int, n-(1<<i)+1)
		prev := data[i-1]
		curr := data[i]
		offset := 1 << (i - 1)
		for j := 0; j <= n-(1<<i); j++ {
			curr[j] = op(prev[j], prev[j+offset])
		}
	}
	return &SparseTable{data: data, op: op}
}

func (st *SparseTable) Query(l, r int) int {
	len := r - l + 1
	k := bits.Len(uint(len)) - 1
	return st.op(st.data[k][l], st.data[k][r-(1<<k)+1])
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
