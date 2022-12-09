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

func testRopes(n int, ps ...pos) (r rope) {
	r = newRope(n)
	for i, p := range ps {
		r.Knots[i].Pos = p
	}
	return
}
func TestMoveHead(t *testing.T) {
	tests := []struct {
		dir  string
		got  rope
		want rope
	}{
		{
			"R",
			testRopes(1, pos{1, 0}, pos{0, 0}),
			testRopes(1, pos{2, 0}, pos{1, 0}),
		},
		{
			"U",
			testRopes(1, pos{4, 0}, pos{3, 0}),
			testRopes(1, pos{4, 1}, pos{3, 0}),
		},
		{
			"U",
			testRopes(1, pos{4, 1}, pos{3, 0}),
			testRopes(1, pos{4, 2}, pos{4, 1}),
		},
		{
			"L",
			testRopes(1, pos{4, 1}, pos{3, 0}),
			testRopes(1, pos{3, 1}, pos{3, 0}),
		},
		{
			"D",
			testRopes(1, pos{4, 1}, pos{3, 0}),
			testRopes(1, pos{4, 0}, pos{3, 0}),
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			tc.got.moveHead(tc.dir)
			for i, knot := range tc.got.Knots {
				got := knot.Pos
				want := tc.want.Knots[i].Pos
				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("knot %d mismatch\n%s", i, diff)
				}
			}
		})
	}
}

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
