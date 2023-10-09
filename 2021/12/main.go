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
	dirs  = map[byte]pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
)

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type pos struct {
	X, Y int
}

type cave struct {
	Id          string
	Connections map[string]struct{}
}

type caves map[string]cave

func parse(lines []string) caves {
	caves := caves{}
	for _, line := range lines {
		start, end, _ := strings.Cut(line, "-")
		if _, ok := caves[start]; !ok {
			caves[start] = cave{
				Id:          start,
				Connections: map[string]struct{}{},
			}
		}
		if _, ok := caves[end]; !ok {
			caves[end] = cave{
				Id:          end,
				Connections: map[string]struct{}{},
			}
		}
		caves[start].Connections[end] = struct{}{}
		caves[end].Connections[start] = struct{}{}
	}
	return caves
}

func dfs(cs caves, c cave, visited map[string]struct{}) (total int) {
	if c.Id == "end" {
		return 1
	}
	for c := range c.Connections {
		if _, ok := visited[c]; !ok {
			// we haven't visited this cave
			if strings.ToLower(c) == c {
				// we shouldn't visit lowercase caves
				visited[c] = struct{}{}
			}
			total += dfs(cs, cs[c], visited)
			delete(visited, c)
		}
	}
	return
}

func part1(lines []string) (res int) {
	caves := parse(lines)
	visited := map[string]struct{}{"start": {}}

	res += dfs(caves, caves["start"], visited)
	return
}

type node struct {
	cave      cave
	seenSmall bool
	path      []string
	visited   map[string]bool
}

func part2(lines []string) (res int) {
	caves := parse(lines)

	stack := []node{{caves["start"], false, []string{"start"}, map[string]bool{}}}
	for len(stack) > 0 {
		// pop the last element
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// continue if we've already visited too many times
		v, ok := cur.visited[cur.cave.Id]
		if v {
			continue
		}

		// copy the history for the new node
		visited := map[string]bool{}
		for k, v := range cur.visited {
			visited[k] = v
		}

		// add the current node to the new history
		if strings.ToLower(cur.cave.Id) == cur.cave.Id {
			if cur.seenSmall || cur.cave.Id == "start" {
				visited[cur.cave.Id] = true
			} else if ok {
				cur.seenSmall = true
				for k := range visited {
					visited[k] = true
				}
			} else {
				visited[cur.cave.Id] = false
			}
		}

		// for each child node
		for c := range caves[cur.cave.Id].Connections {
			// stop if we've hit the end!
			if c == "end" {
				res++
				continue
			}

			// check if we haven't visited too many times
			if v := visited[c]; !v {
				// add the new node to the processing stack
				path := append(cur.path, c)
				stack = append(stack, node{caves[c], cur.seenSmall, path, visited})
			}
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
