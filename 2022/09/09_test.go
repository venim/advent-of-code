package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testInput string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var biggerTestInput string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestGeneric(t *testing.T) {
	tests := []struct {
	}{}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := tc
			want := tc
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

// func TestMoveHead(t *testing.T) {
// 	tests := []struct {
// 		r   rope
// 		dir string
// 		res rope
// 	}{
// 		{
// 			rope{knot{Pos: pos{1, 0}}, knot{Pos: pos{0, 0}, Places: make(map[[2]int]struct{})}},
// 			"R",
// 			rope{knot{Pos: pos{2, 0}}, knot{Pos: pos{1, 0}}},
// 		},
// 		{
// 			rope{knot{Pos: pos{4, 0}}, knot{Pos: pos{3, 0}, Places: make(map[[2]int]struct{})}},
// 			"U",
// 			rope{knot{Pos: pos{4, 1}}, knot{Pos: pos{3, 0}}},
// 		},
// 		{
// 			rope{knot{Pos: pos{4, 1}}, knot{Pos: pos{3, 0}, Places: make(map[[2]int]struct{})}},
// 			"U",
// 			rope{knot{Pos: pos{4, 2}}, knot{Pos: pos{4, 1}}},
// 		},
// 		{
// 			rope{knot{Pos: pos{4, 1}}, knot{Pos: pos{3, 0}, Places: make(map[[2]int]struct{})}},
// 			"L",
// 			rope{knot{Pos: pos{3, 1}}, knot{Pos: pos{3, 0}}},
// 		},
// 		{
// 			rope{knot{Pos: pos{4, 1}}, knot{Pos: pos{3, 0}, Places: make(map[[2]int]struct{})}},
// 			"D",
// 			rope{knot{Pos: pos{4, 0}}, knot{Pos: pos{3, 0}}},
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run("", func(t *testing.T) {
// 			tc.r.moveHead(tc.dir)
// 			if diff := compareHead(tc.r, tc.res); diff != "" {
// 				t.Error(diff)
// 			}
// 			if diff := compareTail(tc.r, tc.res); diff != "" {
// 				t.Error(diff)
// 			}
// 		})
// 	}
// }

// func compareHead(r1, r2 rope) string {
// 	got := r1.h.Pos
// 	want := r2.h.Pos
// 	return cmp.Diff(want, got)
// }

// func compareTail(r1, r2 rope) string {
// 	got := r1.t.Pos
// 	want := r2.t.Pos
// 	return cmp.Diff(want, got)
// }

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{strings.Split(testInput, "\n"), 13},
		{strings.Split(input, "\n"), 5878},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := part1(tc.input)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{strings.Split(biggerTestInput, "\n"), 36},
		{strings.Split(input, "\n"), 2405},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := part2(tc.input)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
