package main

import (
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

	slices.SortFunc(distances, func(a, b Distance) int {
		return a.sq - b.sq
	})

	return jboxes, distances
}

func run(jboxes []*JunctionBox, distances []Distance, cb func(step int, numComponents int, getSizes func() []int, getMergedGroup func() ([]int, int)) (int, bool)) int {
	parent := make([]int, len(jboxes))
	nodes := make([][]int, len(jboxes))
	for i := range jboxes {
		parent[i] = i
		nodes[i] = []int{i}
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

	numComponents := len(jboxes)

	for step, d := range distances {
		rootU := find(d.u)
		rootV := find(d.v)

		var lastMergedNodes []int
		var bridgeNode int

		if rootU != rootV {
			var smallRoot, largeRoot int
			var largeEndpoint int

			if len(nodes[rootU]) < len(nodes[rootV]) {
				smallRoot, largeRoot = rootU, rootV
				largeEndpoint = d.v
			} else {
				smallRoot, largeRoot = rootV, rootU
				largeEndpoint = d.u
			}

			parent[smallRoot] = largeRoot

			// Capture data for Part 2 before merging
			lastMergedNodes = nodes[smallRoot]
			bridgeNode = largeEndpoint

			nodes[largeRoot] = append(nodes[largeRoot], nodes[smallRoot]...)
			nodes[smallRoot] = nil // Release memory

			numComponents--
		}

		if res, done := cb(step, numComponents, func() []int {
			sizes := make([]int, 0, numComponents)
			for i, p := range parent {
				if p == i && len(nodes[i]) > 0 {
					sizes = append(sizes, len(nodes[i]))
				}
			}
			return sizes
		}, func() ([]int, int) {
			return lastMergedNodes, bridgeNode
		}); done {
			return res
		}
	}
	return 0
}

func part1(lines []string, nTimes int) (res int) {
	jboxes, distances := prepare(lines)
	return run(jboxes, distances, func(step int, numComponents int, getSizes func() []int, getMergedGroup func() ([]int, int)) (int, bool) {
		if step == nTimes-1 {
			counts := getSizes()
			slices.Sort(counts)
			res = 1
			start := len(counts) - 3
			if start < 0 {
				start = 0
			}
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
	return run(jboxes, distances, func(step int, numComponents int, getSizes func() []int, getMergedGroup func() ([]int, int)) (int, bool) {
		if numComponents == 1 {
			mergedNodes, bridge := getMergedGroup()
			// Should always have a merge if we just hit 1 component
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
