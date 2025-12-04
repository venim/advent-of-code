package main

import (
	_ "embed"
	"flag"
	"sort"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func parseRanges(lines []string) (idx int, merged []util.Pos) {
	// parse all ranges
	ranges := []util.Pos{}
	i := 0
	for lines[i] != "" {
		parts := strings.Split(lines[i], "-")
		ranges = append(ranges, util.Pos{X: util.MustAtoi(parts[0]), Y: util.MustAtoi(parts[1])})
		i++
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].X < ranges[j].X
	})

	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
		for _, r := range ranges[1:] {
			last := &merged[len(merged)-1]
			if r.X <= last.Y {
				if r.Y > last.Y {
					last.Y = r.Y
				}
			} else {
				merged = append(merged, r)
			}
		}
	}
	return i, merged
}

func part1(lines []string) (res int) {
	i, ranges := parseRanges(lines)

	i++
	for _, available := range lines[i:] {
		n := util.MustAtoi(available)
		idx := sort.Search(len(ranges), func(i int) bool {
			return ranges[i].Y >= n
		})
		if idx < len(ranges) && ranges[idx].X <= n {
			res++
		}
	}

	return
}

func part2(lines []string) (res int) {
	// parse all ranges
	_, ranges := parseRanges(lines)

	// count values in each range
	for _, r := range ranges {
		res += r.Y - r.X + 1
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
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
