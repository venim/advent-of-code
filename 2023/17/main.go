package main

import (
	"container/heap"
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

type grid struct {
	points map[complex128]int
	x, y   int
}

func makeGrid(lines []string) grid {
	x := len(lines[0])
	y := len(lines)
	g := grid{points: make(map[complex128]int), x: x, y: y}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			g.points[complex(float64(j), float64(i))] = util.MustAtoi(string(lines[i][j]))
		}
	}
	return g
}

type node struct {
	pos complex128
	dir complex128
	val int
}

func (n1 node) Less(n2 node) bool {
	return n1.val < n2.val
}

type minHeap []node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = node{}
	*h = old[0 : n-1]
	return x
}

func solve(G grid, min int, max int, end complex128) int {
	q := make(minHeap, 0)
	heap.Init(&q)
	heap.Push(&q, node{0, 1, 0})
	heap.Push(&q, node{0, 1i, 0})

	cache := map[node]bool{}

	for len(q) > 0 {
		cur := heap.Pop(&q).(node)
		if cur.pos == end {
			return cur.val
		}
		cached := node{pos: cur.pos, dir: cur.dir}
		if cache[cached] {
			continue
		}
		cache[cached] = true

		for _, d := range []complex128{1i / cur.dir, -1i / cur.dir} {
			for i := min; i <= max; i++ {
				newPos := cur.pos + d*complex(float64(i), 0)
				if _, ok := G.points[newPos]; ok {
					val := cur.val
					for j := 1; j <= i; j++ {
						val += G.points[cur.pos+d*complex(float64(j), 0)]
					}
					n := node{newPos, d, val}
					heap.Push(&q, n)
				}
			}
		}
	}
	return 0
}

func part1(lines []string) (res int) {
	g := makeGrid(lines)
	end := complex(float64(g.x-1), float64(g.y-1))
	return solve(g, 1, 3, end)
}

func part2(lines []string) (res int) {
	g := makeGrid(lines)
	end := complex(float64(g.x-1), float64(g.y-1))
	return solve(g, 4, 10, end)
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
