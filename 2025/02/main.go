package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

var idRe = regexp.MustCompile(`(\d+)\-(\d+)`)

func isValidPart1(id string) bool {
	n := len(id) / 2
	return id[:n] != id[n:]
}

func part1(lines []string) (res int) {
	for r := range strings.SplitSeq(lines[0], ",") {
		ids := idRe.FindStringSubmatch(r)
		if ids == nil {
			return
		}
		for i := util.MustAtoi(ids[1]); i <= util.MustAtoi(ids[2]); i++ {
			if !isValidPart1(strconv.Itoa(i)) {
				res += i
			}
		}
	}
	return
}

func isValidPart2(id string) bool {
	n := len(id)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			if id == strings.Repeat(id[:i], n/i) {
				return false
			}
		}
	}
	return true
}

func part2(lines []string) (res int) {
	for r := range strings.SplitSeq(lines[0], ",") {
		ids := idRe.FindStringSubmatch(r)
		if ids == nil {
			return
		}
		for i := util.MustAtoi(ids[1]); i <= util.MustAtoi(ids[2]); i++ {
			if !isValidPart2(strconv.Itoa(i)) {
				res += i
			}
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
