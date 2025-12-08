package main

import (
	_ "embed"
	"flag"
	"math"
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
	X       int
	Y       int
	Z       int
	Circuit int
}

func NewJunctionBox(in string) *JunctionBox {
	parts := strings.Split(in, ",")
	return &JunctionBox{
		X: util.MustAtoi(parts[0]),
		Y: util.MustAtoi(parts[1]),
		Z: util.MustAtoi(parts[2]),
	}
}

func (a *JunctionBox) distance(b *JunctionBox) float64 {
	return math.Sqrt(
		math.Pow(float64(b.X-a.X), 2) +
			math.Pow(float64(b.Y-a.Y), 2) +
			math.Pow(float64(b.Z-a.Z), 2))
}

type Distance struct {
	n   float64
	jbs map[*JunctionBox]bool
}

func part1(lines []string, nTimes int) (res int) {
	var jboxes []*JunctionBox
	for _, l := range lines {
		jboxes = append(jboxes, NewJunctionBox(l))
	}

	var distances = []Distance{}
	for i := range jboxes {
		for j := i + 1; j < len(jboxes); j++ {
			start, end := jboxes[i], jboxes[j]
			n := start.distance(end)
			distances = append(distances, Distance{
				n: n,
				jbs: map[*JunctionBox]bool{
					start: true,
					end:   true,
				},
			})
		}
	}

	slices.SortStableFunc(distances, func(a, b Distance) int {
		if a.n < b.n {
			return -1
		}
		if a.n > b.n {
			return 1
		}
		return 0
	})

	nextCircuit := 1
	circuits := map[int]map[*JunctionBox]bool{}
	for range nTimes {
		d := distances[0]
		distances = distances[1:]

		activeCircuit := 0
		// see if any jbox is connected
		connected := false
		for jb := range d.jbs {
			if jb.Circuit != 0 {
				if connected && jb.Circuit != activeCircuit {
					// merge...
					for jb := range circuits[jb.Circuit] {
						d.jbs[jb] = true
					}
					delete(circuits, jb.Circuit)
				} else {
					activeCircuit = jb.Circuit
					connected = true
				}
			}
		}
		// if not connected, create a new circuit
		if activeCircuit == 0 {
			activeCircuit = nextCircuit
			nextCircuit++
			circuits[activeCircuit] = map[*JunctionBox]bool{}
		}
		// set the circuit for all jboxes
		for jb := range d.jbs {
			jb.Circuit = activeCircuit
			circuits[activeCircuit][jb] = true
		}
	}

	var counts []int
	for c := range circuits {
		counts = append(counts, len(circuits[c]))
	}

	slices.Sort(counts)
	res = 1
	for i := len(counts) - 1; i > len(counts)-4; i-- {
		res *= counts[i]
	}

	return
}

func part2(lines []string) (res int) {
	var jboxes []*JunctionBox
	for _, l := range lines {
		jboxes = append(jboxes, NewJunctionBox(l))
	}

	var distances = []Distance{}
	for i := range jboxes {
		for j := i + 1; j < len(jboxes); j++ {
			start, end := jboxes[i], jboxes[j]
			n := start.distance(end)
			distances = append(distances, Distance{
				n: n,
				jbs: map[*JunctionBox]bool{
					start: true,
					end:   true,
				},
			})
		}
	}

	slices.SortStableFunc(distances, func(a, b Distance) int {
		if a.n < b.n {
			return -1
		}
		if a.n > b.n {
			return 1
		}
		return 0
	})

	nextCircuit := 1
	circuits := map[int]map[*JunctionBox]bool{}
	for len(distances) > 0 {
		d := distances[0]
		distances = distances[1:]

		activeCircuit := 0
		// see if any jbox is connected
		connected := false
		for jb := range d.jbs {
			if jb.Circuit != 0 {
				if connected && jb.Circuit != activeCircuit {
					// merge...
					for jb := range circuits[jb.Circuit] {
						d.jbs[jb] = true
					}
					delete(circuits, jb.Circuit)
				} else {
					activeCircuit = jb.Circuit
					connected = true
				}
			}
		}
		// if not connected, create a new circuit
		if activeCircuit == 0 {
			activeCircuit = nextCircuit
			nextCircuit++
			circuits[activeCircuit] = map[*JunctionBox]bool{}
		}
		// set the circuit for all jboxes
		for jb := range d.jbs {
			jb.Circuit = activeCircuit
			circuits[activeCircuit][jb] = true
		}
		if len(circuits) == 1 && len(circuits[activeCircuit]) == len(jboxes) {
			res = 1
			for jb := range d.jbs {
				res *= jb.X
			}
			return
		}
	}
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
	res = part1(lines, 1000)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
