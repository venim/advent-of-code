package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string

	mulRe = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
)

func sumOfMuls(line string) (res int) {
	m := mulRe.FindAllStringSubmatch(line, -1)
	for _, m := range m {
		res += util.MustAtoi(m[1]) * util.MustAtoi(m[2])
	}
	return
}

func part1(line string) (res int) {
	return sumOfMuls(line)
}

func findEnabledChunks(line string) []string {
	chunks := []string{}
	valid := true
	for {
		if valid {
			d := strings.Index(line, "don't()")
			if d != -1 {
				chunks = append(chunks, line[:d])
				line = line[d+7:]
				valid = !valid
				continue
			}
			// the entire line is valid
			chunks = append(chunks, line)
		} else {
			d := strings.Index(line, "do()")
			if d != -1 {
				line = line[d+4:]
				valid = !valid
				continue
			}
		}
		break
	}
	return chunks
}

func part2(line string) (res int) {
	for _, chunk := range findEnabledChunks(line) {
		res += sumOfMuls(chunk)
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

	t = time.Now()
	res = part1(input)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(input)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
