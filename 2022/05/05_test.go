package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testBoard = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestParseBoard(t *testing.T) {
	tests := []struct {
		want [][]string
	}{
		{
			[][]string{
				{"Z", "N"},
				{"M", "C", "D"},
				{"P"},
			},
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			_, got := parseBoard(strings.Split(testBoard, "\n"))
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestParseMove(t *testing.T) {

	tests := []struct {
		line        string
		n, from, to int
	}{
		{"move 1 from 2 to 1", 1, 2, 1},
		{"move 3 from 1 to 3", 3, 1, 3},
		{"move 2 from 2 to 1", 2, 2, 1},
		{"move 1 from 1 to 2", 1, 1, 2},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			gotN, gotFrom, gotTo := parseMove(tc.line)
			got := []int{gotN, gotFrom, gotTo}
			want := []int{tc.n, tc.from, tc.to}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	lines := strings.Split(testBoard, "\n")
	res := part1(lines)
	if res != "CMZ" {
		t.Errorf("%q\n", res)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(testBoard, "\n")
	res := part2(lines)
	if res != "MCD" {
		t.Errorf("%q\n", res)
	}
}
