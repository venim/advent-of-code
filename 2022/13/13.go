package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isList(v any) bool {
	switch v.(type) {
	case []any:
		return true
	}
	return false
}

func makeList(v any) []any {
	switch _v := v.(type) {
	case []any:
		return _v
	default:
		return []any{_v}
	}
}

func compare(left any, right any) int {
	if !isList(left) && !isList(right) {
		if left.(float64) < right.(float64) {
			// left side is smaller
			return -1
		} else if left.(float64) > right.(float64) {
			// right side is smaller
			return 1
		} else {
			return 0
		}
	}
	l := makeList(left)
	r := makeList(right)
	for i := 0; i < len(l); i++ {
		if i >= len(r) {
			// right ran out of items
			return 1
		}
		if diff := compare(l[i], r[i]); diff != 0 {
			// sides don't match
			return diff
		}
	}
	if len(l) < len(r) {
		// left ran out of items
		return -1
	}
	return 0
}

func part1(lines []string) (res int) {
	for i, packets := range lines {
		parts := strings.Split(packets, "\n")
		var left, right []any
		json.Unmarshal([]byte(parts[0]), &left)
		json.Unmarshal([]byte(parts[1]), &right)
		if compare(left, right) == -1 {
			res += (i + 1)
		}
	}
	return
}

func part2(lines []string) (res int) {
	res = 1
	packets := make([]any, 0, len(lines)*2)
	dividers := []string{"[[2]]", "[[6]]"}
	lines = append(lines, dividers...)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var packet []any
		json.Unmarshal([]byte(line), &packet)
		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == -1
	})
	for i, p := range packets {
		for _, d := range dividers {
			if fmt.Sprint(p) == d {
				res *= (i + 1)
			}
		}
	}
	return
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t     time.Time
		res   int
		lines []string
	)

	lines = strings.Split(input, "\n\n")
	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	lines = strings.Split(input, "\n")
	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
