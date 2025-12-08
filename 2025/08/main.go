package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

type JunctionBox struct {
	X int
	Y int
	Z int
}

func NewJunctionBox(in string) *JunctionBox {
	parts := strings.Split(in, ",")
	return &JunctionBox{
		X: util.MustAtoi(parts[0]),
		Y: util.MustAtoi(parts[1]),
		Z: util.MustAtoi(parts[2]),
	}
}

func (a *JunctionBox) distSq(b *JunctionBox) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return dx*dx + dy*dy + dz*dz
}

type Distance struct {
	u, v int
	sq   int
}

// Less method for Distance to satisfy util.MinHeapItem interface
func (d Distance) Less(other Distance) bool {
	return d.sq < other.sq
}

func prepare(lines []string) ([]*JunctionBox, []Distance) {
	var jboxes []*JunctionBox
	for _, l := range lines {
		jboxes = append(jboxes, NewJunctionBox(l))
	}

	n := len(jboxes)
	distances := make([]Distance, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distances = append(distances, Distance{
				u:  i,
				v:  j,
				sq: jboxes[i].distSq(jboxes[j]),
			})
		}
	}
	// Return unsorted distances. The heap will handle ordering.
	return jboxes, distances
}

func run(jboxes []*JunctionBox, allDistances []Distance, trackNodes bool, cb func(step int, numComponents int, sizes []int, getMergedGroup func() ([]int, int)) (int, bool)) int {
	numJboxes := len(jboxes)
	parent := make([]int, numJboxes)
	sizes := make([]int, numJboxes)

	var nodes [][]int
	if trackNodes {
		nodes = make([][]int, numJboxes)
		for i := range jboxes {
			nodes[i] = []int{i}
		}
	}

	for i := range jboxes {
		parent[i] = i
		sizes[i] = 1
	}

	find := func(i int) int {
		root := i
		for parent[root] != root {
			root = parent[root]
		}
		curr := i
		for curr != root {
			next := parent[curr]
			parent[curr] = root
			curr = next
		}
		return root
	}

	// Initialize Heap using util.MinHeap
	// Copy distances to avoid modifying the original slice shared between parts
	pq := make(util.MinHeap[Distance], len(allDistances))
	copy(pq, allDistances)
	heap.Init(&pq)

	numComponents := numJboxes
	step := 0

	for pq.Len() > 0 {
		d := heap.Pop(&pq).(Distance)

		rootU := find(d.u)
		rootV := find(d.v)

		var lastMergedNodes []int
		var bridgeNode int

		if rootU != rootV {
			var smallRoot, largeRoot int
			var largeEndpoint int

			// Union by size
			if sizes[rootU] < sizes[rootV] {
				smallRoot, largeRoot = rootU, rootV
				largeEndpoint = d.v
			} else {
				smallRoot, largeRoot = rootV, rootU
				largeEndpoint = d.u
			}

			parent[smallRoot] = largeRoot
			sizes[largeRoot] += sizes[smallRoot]
			sizes[smallRoot] = 0 // The smaller component is now part of the larger one
			// Only maintain node lists if required
			if trackNodes {
				lastMergedNodes = nodes[smallRoot]
				bridgeNode = largeEndpoint

				nodes[largeRoot] = append(nodes[largeRoot], nodes[smallRoot]...)
				nodes[smallRoot] = nil // Release memory
			}

			numComponents--
		}

		if res, done := cb(step, numComponents, sizes, func() ([]int, int) {
			return lastMergedNodes, bridgeNode
		}); done {
			return res
		}
		step++
	}
	return 0
}

func part1(lines []string, nTimes int) (res int) {
	jboxes, distances := prepare(lines)
	// trackNodes is false for Part 1 as it doesn't need to know the specific nodes in each component, saving memory and time.
	return run(jboxes, distances, false, func(step int, numComponents int, sizes []int, getMergedGroup func() ([]int, int)) (int, bool) {
		if step == nTimes-1 { // Stop after nTimes connections
			var counts []int
			// Collect sizes of active components (those with size > 0, indicating they are roots)
			for _, s := range sizes {
				if s > 0 {
					counts = append(counts, s)
				}
			}
			slices.Sort(counts) // Sort component sizes to find the largest ones
			res = 1
			// Multiply the sizes of the 3 largest components
			start := max(len(counts)-3, 0)
			for i := len(counts) - 1; i >= start; i-- {
				res *= counts[i]
			}
			return res, true
		}
		return 0, false
	})
}

func part2(lines []string) (res int) {
	jboxes, distances := prepare(lines)
	// trackNodes=true needed for logic
	return run(jboxes, distances, true, func(step int, numComponents int, sizes []int, getMergedGroup func() ([]int, int)) (int, bool) {
		if numComponents == 1 {
			mergedNodes, bridge := getMergedGroup()
			if mergedNodes != nil {
				res = 1
				for _, idx := range mergedNodes {
					res *= jboxes[idx].X
				}
				res *= jboxes[bridge].X
				return res, true
			}
		}
		return 0, false
	})
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
	res = part1(lines, 1000)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
