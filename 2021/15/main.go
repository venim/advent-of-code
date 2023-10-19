package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input  string
	dirs   = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

func (p pos) distanceTo(goal pos) int {
	dx := math.Abs(float64(p.X - goal.X))
	dy := math.Abs(float64(p.Y - goal.Y))
	return int(dx + dy)
}

type node struct {
	p      pos
	fScore int
}

type minHeap []*node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].fScore < h[j].fScore }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*node))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

func parse(lines []string, n int) [][]int {
	x := len(lines)
	y := len(lines[0])
	grid := make([][]int, 0, x)
	for i := 0; i < x*n; i++ {
		grid = append(grid, make([]int, 0, y))
		for j := 0; j < y*n; j++ {
			n := mustAtoi(string(lines[i%x][j%y]))
			n += i/x + j/y
			if n > 9 {
				n %= 9
			}
			grid[i] = append(grid[i], n)
		}
	}
	return grid
}

func aStarPriorityQueue(grid [][]int) (res int) {
	start := pos{}
	end := pos{len(grid) - 1, len(grid[0]) - 1}

	openSet := make(minHeap, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &node{start, start.distanceTo(end)})
	cameFrom := map[pos]pos{}

	gScore := map[pos]int{}
	gScore[start] = 0

	for {
		if len(openSet) == 0 {
			return
		}

		current := heap.Pop(&openSet).(*node).p

		if current == end {
			for {
				res += grid[current.X][current.Y]
				next, ok := cameFrom[current]
				if ok && next == start {
					return res
				}
				current = next
			}
		}

		for _, d := range dirs {
			neighbor := pos{current.X + d.X, current.Y + d.Y}
			if neighbor.X >= 0 && neighbor.X <= end.X &&
				neighbor.Y >= 0 && neighbor.Y <= end.Y {
				newGScore := gScore[current] + grid[neighbor.X][neighbor.Y]
				if oldGScore, ok := gScore[neighbor]; !ok || newGScore < oldGScore {
					cameFrom[neighbor] = current
					gScore[neighbor] = newGScore
					heap.Push(&openSet, &node{neighbor, newGScore + neighbor.distanceTo(end)})
				}
			}
		}
	}
}

func part1(lines []string) (res int) {
	grid := parse(lines, 1)

	return aStarPriorityQueue(grid)
}
func part2(lines []string) (res int) {
	grid := parse(lines, 5)

	return aStarPriorityQueue(grid)
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
