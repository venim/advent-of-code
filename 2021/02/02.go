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

type pos struct {
	X, Y, Z int
}

func part1(lines []string) (res int) {
	p := new(pos)
	for _, line := range lines {
		command := strings.Split(line, " ")
		dir := command[0]
		n := mustAtoi(command[1])
		switch dir {
		case "forward":
			p.X += n
		case "up":
			p.Y -= n
		case "down":
			p.Y += n
		}
	}
	return p.X * p.Y
}

func part2(lines []string) (res int) {
	p := new(pos)
	for _, line := range lines {
		command := strings.Split(line, " ")
		dir := command[0]
		n := mustAtoi(command[1])
		switch dir {
		case "forward":
			// increases your horizontal position by X units.
			p.X += n
			// increases your depth by your aim multiplied by X.
			p.Y += p.Z * n
		case "up":
			// decreases your aim by X units.
			p.Z -= n
		case "down":
			// increases your aim by X units.
			p.Z += n
		}
	}
	return p.X * p.Y
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
