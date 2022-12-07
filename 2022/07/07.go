package main

import (
	_ "embed"
	"fmt"
	"sort"

	"strconv"
	"strings"
	"time"
)

var (
	//go:embed in.txt
	input string
)

type directory struct {
	parent *directory
	files  map[string]int
	dirs   map[string]*directory
	size   int
}

func newDir(parent *directory) *directory {
	return &directory{
		parent: parent,
		files:  make(map[string]int),
		dirs:   make(map[string]*directory),
	}
}

func (d *directory) du() {
	for _, n := range d.files {
		d.size += n
	}
	for _, dir := range d.dirs {
		dir.du()
		d.size += dir.size
	}
}

type fileSystem struct {
	commands []string
	root     *directory
	cwd      *directory
}

func (f *fileSystem) run(commands []string) {
	for _, cmd := range commands {
		if strings.Index(cmd, "$ ls") == 0 {
			continue
		} else if strings.Index(cmd, "$ cd") == 0 {
			dir := strings.Split(cmd, " ")[2]
			f.cd(dir)
		} else {
			f.touch(cmd)
		}
	}
}

func (f *fileSystem) cd(dir string) {
	switch dir {
	case "..":
		f.cwd = f.cwd.parent
	case "/":
		f.root = newDir(nil)
		f.cwd = f.root
	default:
		d := newDir(f.cwd)
		f.cwd.dirs[dir] = d
		f.cwd = d
	}
}

func (f *fileSystem) touch(cmd string) {
	line := strings.Split(cmd, " ")
	if line[0] == "dir" {
		f.cwd.dirs[line[1]] = newDir(f.cwd)
	} else {
		f.cwd.files[line[1]] = mustAtoi(line[0])
	}
}

func sum(d *directory, maxSize int) (total int) {
	if d.size <= maxSize {
		total += d.size
	}
	for _, dir := range d.dirs {
		total += sum(dir, maxSize)
	}
	return total
}

func dirSizes(d *directory) (size []int) {
	size = append(size, d.size)
	for _, dir := range d.dirs {
		size = append(size, dirSizes(dir)...)
	}
	return
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func part1(lines []string) (res int) {
	f := fileSystem{}
	f.run(lines)
	f.root.du()
	res = sum(f.root, 100000)
	return res
}

func part2(lines []string) (res int) {
	f := fileSystem{}
	f.run(lines)
	f.root.du()
	over := f.root.size - 40000000
	sizes := dirSizes(f.root)
	sort.Ints(sizes)
	for _, n := range sizes {
		if n >= over {
			return n
		}
	}
	return
}

func main() {
	var (
		t   time.Time
		res int
	)
	lines := strings.Split(input, "\n")

	t = time.Now()
	res = part1(lines)
	fmt.Printf("[Part 1] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))

	t = time.Now()
	res = part2(lines)
	fmt.Printf("[Part 2] = %v\n", res)
	fmt.Printf("took %s\n\n", time.Since(t))
}
