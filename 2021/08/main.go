package main

import (
	_ "embed"
	"flag"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		digits := strings.Fields(strings.Split(line, " | ")[1])
		for _, d := range digits {
			switch len(d) {
			case 2:
				// it's a 1
				res++
			case 3:
				// it's a 7
				res++
			case 4:
				// it's a 4
				res++
			case 7:
				// it's an 8
				res++
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		segments := make([]rune, 7)
		knownDigits := map[string]string{}
		digits := strings.Fields(line)
		for _, d := range digits {
			switch len(d) {
			case 2:
				// it's a 1
				knownDigits["1"] = d
			case 3:
				// it's a 7
				knownDigits["7"] = d
			case 4:
				// it's a 4
				knownDigits["4"] = d
			case 7:
				// it's an 8
				knownDigits["8"] = d
			}
		}
		// segment a exists in 7, but not 1
		segments[0] = rune(difference(knownDigits["7"], (knownDigits["1"]))[0])

		// segment c exists in 1, not 6
		for _, d := range digits {
			// digits 0, 6, 9 have 6 segments
			if len(d) == 6 {
				// only 6 doesn't include segment c
				if i := difference(knownDigits["1"], d); len(i) == 1 {
					knownDigits["6"] = d
					segments[2] = i[0]
					break
				}
			}
		}

		// segment f is the other one in 1
		for _, s := range knownDigits["1"] {
			if s != segments[2] {
				segments[5] = s
				break
			}
		}

		// segment b exists in 4, not 3 (and we know segments c & f)
		for _, d := range digits {
			// digits 3, 2, 5 have 5 segments
			if len(d) == 5 {
				// digit 3 has 2 segments left after removing segments a, c, & f
				remaining := make([]rune, 0, 3)
				for _, s := range d {
					if !(s == segments[2] ||
						s == segments[0] ||
						s == segments[5]) {
						remaining = append(remaining, s)
					}
				}
				if len(remaining) == 2 {
					knownDigits["3"] = d
					segments[1] = rune(difference(knownDigits["4"], d)[0])
					break
				}
			}
		}

		// segment d is left after removing segments b, c, f from 4
		segments[3] = rune(difference(knownDigits["4"], string([]rune{segments[1], segments[2], segments[5]}))[0])

		// segment g is left after removing segments a, c, d, f from 3
		segments[6] = rune(difference(knownDigits["3"], string([]rune{segments[0], segments[2], segments[3], segments[5]}))[0])

		// segment e is left after removing segments a, b, c, d, f, g from 8
		segments[4] = rune(difference(knownDigits["8"], string([]rune{segments[0], segments[1], segments[2], segments[3], segments[5], segments[6]}))[0])

		// know the remaining digits
		for _, d := range digits {
			two_segments := []rune{segments[0], segments[2], segments[3], segments[4], segments[6]}
			if _, ok := knownDigits["2"]; !ok && len(d) == 5 {
				intersection := intersect(d, string(two_segments))
				if slices.Equal(two_segments, []rune(intersection)) {
					knownDigits["2"] = d
				}
			}

			five_segments := []rune{segments[0], segments[1], segments[3], segments[5], segments[6]}
			if _, ok := knownDigits["5"]; !ok && len(d) == 5 {
				if slices.Equal(five_segments, []rune(intersect(d, string(five_segments)))) {
					knownDigits["5"] = d
				}
			}

			nine_segments := []rune{segments[0], segments[1], segments[2], segments[3], segments[5], segments[6]}
			if _, ok := knownDigits["9"]; !ok && len(d) == 6 {
				if slices.Equal(nine_segments, []rune(intersect(d, string(nine_segments)))) {
					knownDigits["9"] = d
				}
			}

			zero_segments := []rune{segments[0], segments[1], segments[2], segments[4], segments[5], segments[6]}
			if _, ok := knownDigits["0"]; !ok && len(d) == 6 {
				if slices.Equal(zero_segments, []rune(intersect(d, string(zero_segments)))) {
					knownDigits["0"] = d
				}
			}
		}

		// generate a map of each digit encoding (sorted)
		hash := map[string]string{}
		for d, s := range knownDigits {
			r := []rune(s)
			slices.Sort(r)
			hash[string(r)] = d
		}
		// decode each digit and add to the number
		n := ""
		for _, s := range strings.Fields(strings.Split(line, "|")[1]) {
			r := []rune(s)
			slices.Sort(r)
			n += hash[string(r)]
		}
		// add the number to the total
		res += mustAtoi(n)
	}
	return
}

func intersect(s1, s2 string) []rune {
	set := make([]rune, 0)
	hash := make(map[rune]struct{})

	for _, v := range s1 {
		hash[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}
	return set
}

func difference(a, b string) []rune {
	in := map[rune]struct{}{}
	for _, s := range b {
		in[s] = struct{}{}
	}
	diff := make([]rune, 0)
	for _, s := range a {
		if _, ok := in[s]; !ok {
			diff = append(diff, s)
		}
	}
	return diff
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
