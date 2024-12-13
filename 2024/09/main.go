package main

import (
	_ "embed"
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

type file struct {
	Id    int
	Size  int
	Start int
	End   int
}

func parse(input string) []*file {
	files := []*file{}
	index := 0
	for i, v := range input {
		v := int(v) - 48
		if i%2 == 0 {
			// this is a file
			f := &file{
				Id:    i / 2,
				Size:  v,
				Start: index,
				End:   index + v - 1,
			}
			files = append(files, f)
		} else {
			// this is free space
		}
		index += v
	}
	return files
}

func part1(lines []string) (res int) {
	files := parse(lines[0])
	left := 0
	pos := 0
	right := len(files) - 1
	needed := files[right].Size
	for left < right {
		if needed == 0 {
			right--
			needed = files[right].Size
			continue
		}
		available := files[left].Start - pos
		if available > needed {
			start := pos
			pos += needed
			for i := start; i < pos; i++ {
				res += i * right
			}
			needed = 0
			continue
		}
		needed -= available
		start := pos
		pos = files[left].End + 1
		for i := start; i < start+available; i++ {
			res += i * right
		}
		for i := start + available; i < pos; i++ {
			res += i * left
		}
		left++
	}
	for i := pos; i < pos+needed; i++ {
		res += i * left
	}
	return
}

func moveEntireFile(f *file, files []*file) int {
	pos := 0
	for i := 0; i < f.Id; i++ {
		available := files[i].Start - pos
		if available >= f.Size {
			// "reserve" space occupied after the move
			files[i-1].End += f.Size
			return pos
		}
		pos = files[i].End + 1
	}
	return -1
}

func part2(lines []string) (res int) {
	files := parse(lines[0])
	right := len(files)

	for {
		right--
		if right == 0 {
			return
		}
		f := files[right]
		if start := moveEntireFile(f, files); start != -1 {
			f.Start = start
		}
		// score
		for i := f.Start; i < f.Start+f.Size; i++ {
			res += i * f.Id
		}

	}
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
