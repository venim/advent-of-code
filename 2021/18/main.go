package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"math"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func addLeft(x any, n any) any {
	if n == nil {
		return x
	}
	if v, ok := x.(float64); ok {
		return v + n.(float64)
	}
	v := x.([]any)
	return []any{addLeft(v[0], n), v[1]}
}

func addRight(x any, n any) any {
	if n == nil {
		return x
	}
	if v, ok := x.(float64); ok {
		return v + n.(float64)
	}
	v := x.([]any)
	return []any{v[0], addRight(v[1], n)}
}

func explode(x any, n int) (any, any, any, bool) {
	if v, ok := x.(float64); ok {
		return nil, v, nil, false
	}
	v := x.([]any)
	if n == 0 {
		return v[0], float64(0), v[1], true
	}
	var (
		a, b, l, r any
		ok         bool
	)
	a, b = v[0], v[1]
	l, a, r, ok = explode(a, n-1)
	if ok {
		return l, []any{a, addLeft(b, r)}, nil, true
	}
	l, b, r, ok = explode(b, n-1)
	if ok {
		return nil, []any{addRight(a, l), b}, r, true
	}
	return nil, x, nil, false
}

func split(x any) (any, bool) {
	if v, ok := x.(float64); ok {
		if v >= 10 {
			return []any{math.Floor(v / 2), math.Ceil(v / 2)}, true
		}
		return x, false
	}
	n := x.([]any)
	a, b := n[0], n[1]
	var ok bool
	a, ok = split(a)
	if ok {
		return []any{a, b}, true
	}
	b, ok = split(b)
	return []any{a, b}, ok
}

func add(a, b any) any {
	var (
		x  any = []any{a, b}
		ok bool
	)
	for {
		_, x, _, ok = explode(x, 4)
		if ok {
			continue
		}
		x, ok = split(x)
		if !ok {
			break
		}
	}
	return x
}

func magnitude(x any) int {
	if v, ok := x.(float64); ok {
		return int(v)
	}
	n := x.([]any)
	return 3*magnitude(n[0]) + 2*magnitude(n[1])
}

func part1(lines []string) (res int) {
	sailfishNumbers := []any{}
	for _, l := range lines {
		var d any
		json.Unmarshal([]byte(l), &d)
		sailfishNumbers = append(sailfishNumbers, d)
	}
	root := sailfishNumbers[0]
	for _, n := range sailfishNumbers[1:] {
		root = add(root, n)
	}

	glog.Infof("%v", root)
	return magnitude(root)
}

func part2(lines []string) (res int) {
	sailfishNumbers := []any{}
	for _, l := range lines {
		var d any
		json.Unmarshal([]byte(l), &d)
		sailfishNumbers = append(sailfishNumbers, d)
	}
	combinations := [][]any{}
	for i := 0; i < len(sailfishNumbers); i++ {
		for j := 0; j < len(sailfishNumbers); j++ {
			if i == j {
				continue
			}
			combinations = append(combinations, []any{sailfishNumbers[i], sailfishNumbers[j]})
		}
	}
	for _, c := range combinations {
		res = max(res, magnitude(add(c[0], c[1])))
	}
	return
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
