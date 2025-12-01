package main

import (
	_ "embed"
	"flag"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

func part1(lines []string) (res int) {
	pos := 50
	for _, l := range lines {
		if l == "" {
			break
		}
		n := util.MustAtoi(l[1:])
		switch l[0] {
		case 'L':
			pos -= n
			for pos < 0 {
				pos += 100
			}
		case 'R':
			pos += n
			for pos > 99 {
				pos -= 100
			}
		}
		if pos == 0 {
			res++
		}
	}
	return
}

func part2(lines []string) (res int) {
	pos := 50
	for _, l := range lines {
		if l == "" {
			break
		}
		n := util.MustAtoi(l[1:])
		res += n / 100
		n %= 100
		switch l[0] {
		case 'L':
			pos -= n
			for pos < 0 {
				if pos != -n {
					res++
				}
				pos += 100
			}
		case 'R':
			pos += n
			for pos > 99 {
				if pos != 100 {
					res++
				}
				pos -= 100
			}
		}
		if pos == 0 {
			res++
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
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
