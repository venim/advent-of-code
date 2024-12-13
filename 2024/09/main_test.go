package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in_test.txt
	test string
)

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

func TestParse(t *testing.T) {
	want := []*file{
		{0, 2, 0, 1},
		{1, 3, 5, 7},
		{2, 1, 11, 11},
		{3, 3, 15, 17},
		{4, 2, 19, 20},
		{5, 4, 22, 25},
		{6, 4, 27, 30},
		{7, 3, 32, 34},
		{8, 4, 36, 39},
		{9, 2, 40, 41},
	}
	got := parse(test)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: util.SplitLines(test),
			want:  1928,
		},
		{
			input: util.SplitLines(input),
			want:  6446899523367,
		},
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
		{
			input: util.SplitLines(test),
			want:  2858,
		},
		{
			input: util.SplitLines(input),
			want:  6478232739671,
		},
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
