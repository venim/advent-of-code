package main

import (
	_ "embed"
	"flag"
	"regexp"
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

func hash(line string) (v int) {
	for _, s := range line {
		v += int(s)
		v *= 17
		v %= 256
	}
	return
}

func part1(lines []string) (res int) {
	for _, n := range strings.Split(lines[0], ",") {
		res += hash(n)
	}
	return
}

type lens struct {
	Label       string
	FocalLength int
}

func lensCompare(label string) func(lens) bool {
	return func(l lens) bool {
		return l.Label == label
	}
}

type box struct {
	Lenses []lens
}

func (b *box) add(l lens) {
	i := slices.IndexFunc(b.Lenses, lensCompare(l.Label))
	if i != -1 {
		b.Lenses[i] = l
		return
	}
	b.Lenses = append(b.Lenses, l)
}

func (b *box) remove(label string) {
	b.Lenses = slices.DeleteFunc(b.Lenses, lensCompare(label))
}

var (
	lineRe = regexp.MustCompile(`(\w+)([-=])(.*)?`)
)

func part2(lines []string) (res int) {
	boxes := make([]box, 256)
	for _, l := range strings.Split(lines[0], ",") {
		m := lineRe.FindStringSubmatch(l)
		h := hash(m[1])
		switch m[2] {
		case "=":
			boxes[h].add(lens{m[1], util.MustAtoi(m[3])})
		case "-":
			boxes[h].remove(m[1])
		}
	}
	for i, b := range boxes {
		for j, l := range b.Lenses {
			score := (i + 1) * (j + 1) * l.FocalLength
			res += score
		}
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
