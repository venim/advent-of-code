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

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: util.SplitLines(test),
			want:  2,
		},
		{
			input: util.SplitLines(input),
			want:  585,
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

func TestRemoveLevel(t *testing.T) {
	tests := []struct {
		levels []int
		index  int
		want   []int
	}{
		{[]int{1, 3, 2, 4, 5}, 1, []int{1, 2, 4, 5}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := removeLevel(tc.levels, tc.index)
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
			want:  4,
		},
		{
			input: util.SplitLines(input),
			want:  626,
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
