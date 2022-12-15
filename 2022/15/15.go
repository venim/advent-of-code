package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
	re    = regexp.MustCompile(`.*x=(-?\d+), y=(-?\d+).*x=(-?\d+), y=(-?\d+)`)
)

type point struct {
	X int
	Y int
}

func (p point) dist(p2 point) int {
	dx := p.X - p2.X
	dy := p.Y - p2.Y
	return absInt(dx) + absInt(dy)
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func updateIneligibile(locations map[point]struct{}, s, b point, d, y int) {
	if y >= s.Y-d && y <= s.Y+d {
		x := d - absInt(s.Y-y)

		for i := -x; i < x+1; i++ {
			p := point{s.X + i, y}
			if p == b {
				continue
			}
			locations[p] = struct{}{}
		}
	}
}

func findBeacon(sensors map[point]int, tuning int) int {
	for i := 0; i <= tuning; i++ {
		for j := 0; j <= tuning; j++ {
			beacon := true
			p := point{i, j}
			for s, d := range sensors {
				if s.dist(p) <= d {
					j = s.Y + d - absInt(s.X-i)
					beacon = false
					break
				}
			}
			if beacon {
				return i*4000000 + j
			}
		}
	}
	return 0
}

func solve(lines []string, row int, tuning int) (part1 int, part2 int) {
	locations := map[point]struct{}{}
	sensors := make(map[point]int, len(lines))
	beacons := make(map[point]int, len(lines))
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		s := point{mustAtoi(match[1]), mustAtoi(match[2])}
		b := point{mustAtoi(match[3]), mustAtoi(match[4])}
		d := s.dist(b)
		sensors[s] = d
		beacons[b] = d
		updateIneligibile(locations, s, b, d, row)
		if row >= s.Y-d && row <= s.Y+d {
			x := d - absInt(s.Y-row)

			for i := -x; i < x+1; i++ {
				p := point{s.X + i, row}
				if p == b {
					continue
				}
				locations[p] = struct{}{}
			}
		}
	}
	part1 = len(locations)
	part2 = findBeacon(sensors, tuning)
	return
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t time.Time
	)
	flag.Parse()
	lines := strings.Split(input, "\n")

	t = time.Now()
	part1, part2 := solve(lines, 2000000, 4000000)
	glog.Infof("[Part 1] = %v\n", part1)
	glog.Infof("[Part 2] = %v\n", part2)
	glog.Infof("took %s\n\n", time.Since(t))
}
