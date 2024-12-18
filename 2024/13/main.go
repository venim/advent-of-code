package main

import (
	_ "embed"
	"flag"
	"regexp"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

type pos struct {
	X float64
	Y float64
}

type machine struct {
	A pos
	B pos
	P pos
}

func isWhole(f float64) bool {
	return int(f*100) == int(f)*100
}

func (m *machine) coeff() (a float64, b float64) {
	b = (m.P.Y*m.A.X - m.P.X*m.A.Y) / (m.B.Y*m.A.X - m.B.X*m.A.Y)
	a = (m.P.X - b*m.B.X) / m.A.X

	return
}

var buttonRe = regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
var prizeRe = regexp.MustCompile(`X=(\d+), Y=(\d+)`)

func aToF(s string) float64 {
	return float64(util.MustAtoi(s))
}

func parse(lines []string) []*machine {
	ms := []*machine{}
	for i := 0; i < len(lines); i += 4 {
		aLine := buttonRe.FindStringSubmatch(lines[i])
		bLine := buttonRe.FindStringSubmatch(lines[i+1])
		pLine := prizeRe.FindStringSubmatch(lines[i+2])
		ms = append(ms, &machine{
			A: pos{aToF(aLine[1]), aToF(aLine[2])},
			B: pos{aToF(bLine[1]), aToF(bLine[2])},
			P: pos{aToF(pLine[1]), aToF(pLine[2])},
		})
	}
	return ms
}

func part1(lines []string) (res int) {
	machines := parse(lines)
	for _, m := range machines {
		a, b := m.coeff()
		if a <= 100 && b <= 100 && isWhole(a) && isWhole(b) {
			res += int(3*a + b)
		}
	}
	return
}

func part2(lines []string) (res int) {
	machines := parse(lines)
	for _, m := range machines {
		m.P.X += 10000000000000
		m.P.Y += 10000000000000
		a, b := m.coeff()
		if isWhole(a) && isWhole(b) {
			res += int(3*a + b)
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
	lines := util.SplitLines(input)

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
