package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var test string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

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

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{strings.Split(test, "\n"), 31},
		{strings.Split(input, "\n"), 481},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			grid, start, end, _ := makeGrid(tc.input)
			got := findPath(grid, []node{{start, 0}}, end)
			if diff := cmp.Diff(tc.want, got); diff != "" {
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
		{strings.Split(test, "\n"), 29},
		{strings.Split(input, "\n"), 480},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			grid, _, end, starts := makeGrid(tc.input)
			got := findPath(grid, starts, end)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
