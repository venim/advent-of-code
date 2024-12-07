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

func TestEval(t *testing.T) {
	tests := []struct {
		nums []int
		ops  []byte
		want int
	}{
		{[]int{10, 19}, []byte{'+'}, 29},
		{[]int{10, 19}, []byte{'*'}, 190},
		{[]int{81, 40, 27}, []byte{'+', '*'}, 3267},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := eval(tc.nums, tc.ops)
			want := tc.want
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
			want:  3749,
		},
		{
			input: util.SplitLines(input),
			want:  28730327770375,
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
			want:  11387,
		},
		{
			input: util.SplitLines(input),
			want:  424977609625985,
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
