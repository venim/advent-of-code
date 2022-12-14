package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var test string = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

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

func TestMakeRocks(t *testing.T) {
	tests := []struct {
		input string
		want  []point
	}{
		{"498,4 -> 498,6 -> 496,6", []point{{498, 4}, {498, 5}, {498, 6}, {497, 6}, {496, 6}}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			g := newGrid()
			g.makeRocks(tc.input)
			for _, p := range tc.want {
				if _, ok := g.Rocks[p]; !ok {
					t.Errorf("missing point %+v", p)
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
		{strings.Split(test, "\n"), 24},
		{strings.Split(input, "\n"), 888},
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
		{strings.Split(test, "\n"), 93},
		{strings.Split(input, "\n"), 26461},
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
