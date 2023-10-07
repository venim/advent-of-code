package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			input: strings.Split(test, "\n"),
			want:  5934,
		},
		{
			input: strings.Split(input, "\n"),
			want:  345387,
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
			input: strings.Split(test, "\n"),
			want:  26984457539,
		},
		{
			input: strings.Split(input, "\n"),
			want:  1574445493136,
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

func TestShift(t *testing.T) {
	tests := []struct {
		start []int
		want  []int
	}{
		{
			[]int{0, 1, 1, 2, 1, 0, 0, 0, 0},
			[]int{1, 1, 2, 1, 0, 0, 0, 0, 0},
		},
		{
			// 1 fish with a 0 timer adds to column 6 and 8
			[]int{1, 1, 2, 1, 0, 0, 0, 0, 0},
			[]int{1, 2, 1, 0, 0, 0, 1, 0, 1},
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := shift(tc.start)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
