package main

import (
	_ "embed"
	"flag"
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

type node struct {
	Parent *node
	Child  *node
	V      int
}

type file struct {
	Head *node
	Len  int
}

func (f *file) getAfter(v, n int) int {
	cur := f.Head
	for {
		cur = cur.Child
		if cur.V == v {
			break
		}
	}
	for i := 0; i < n%(f.Len); i++ {
		cur = cur.Child
	}
	return cur.V
}

func (f *file) slice() []*node {
	s := make([]*node, 0, f.Len)
	cur := f.Head
	for i := 0; i < f.Len; i++ {
		s = append(s, cur)
		cur = cur.Child
	}
	return s
}

func (f *file) sliceInt() []int {
	s := make([]int, 0, f.Len)
	cur := f.Head
	for i := 0; i < f.Len; i++ {
		s = append(s, cur.V)
		cur = cur.Child
	}
	return s
}

func (f *file) mix(order []*node) {
	for _, n := range order {
		f.shift(n)
	}
}

func (f *file) rShift(n *node) {
	c := n.Child
	p := n.Parent

	p.Child = c
	c.Parent = p
	c.Child.Parent = n
	n.Child = c.Child
	c.Child = n
	n.Parent = c

	if f.Head == n {
		f.Head = c
	}
}

func (f *file) lShift(n *node) {
	c := n.Child
	p := n.Parent

	c.Parent = p
	p.Child = c
	p.Parent.Child = n
	n.Parent = p.Parent
	p.Parent = n
	n.Child = p

	if f.Head == n {
		f.Head = p
	}
}

func (f *file) shift(n *node) {
	v := n.V % (f.Len - 1)

	if v > 0 {
		for i := 0; i < v; i++ {
			f.rShift(n)
		}
	} else if v < 0 {
		for i := v; i < 0; i++ {
			f.lShift(n)
		}
	}
}

func makeFile(lines []string, key int) *file {
	f := new(file)
	f.Len = len(lines)
	var n, cur *node
	for i, l := range lines {
		n = &node{Parent: cur, V: mustAtoi(l) * key}
		if i == 0 {
			cur = n
			f.Head = n
			continue
		}
		cur.Child = n
		cur = n
	}
	cur.Child = f.Head
	f.Head.Parent = cur
	return f
}

func part1(lines []string) (res int) {
	f := makeFile(lines, 1)
	f.mix(f.slice())
	for _, i := range []int{1000, 2000, 3000} {
		res += f.getAfter(0, i)
	}
	return
}

func part2(lines []string) (res int) {
	f := makeFile(lines, 811589153)
	order := f.slice()
	for i := 0; i < 10; i++ {
		f.mix(order)
	}
	for _, i := range []int{1000, 2000, 3000} {
		res += f.getAfter(0, i)
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
