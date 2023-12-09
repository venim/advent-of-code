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

func TestGenerate(t *testing.T) {
	tests := []struct {
		in      string
		want    history
		reverse bool
	}{
		{
			"0 3 6 9 12 15",
			history{Sequences: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			}},
			false,
		},
		{
			"1 3 6 10 15 21",
			history{Sequences: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			}},
			false,
		},
		{
			"10 13 16 21 30 45",
			history{Sequences: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			}},
			false,
		},
		{
			"-7 -12 -15 -4 33",
			history{Sequences: [][]int{
				{-7, -12, -15, -4, 33},
				{-5, -3, 11, 37},
				{2, 14, 26},
				{12, 12},
				{0},
			}},
			false,
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := generate(tc.in, tc.reverse)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestExtrapolate(t *testing.T) {
	tests := []struct {
		in   history
		want history
	}{
		{
			history{Sequences: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			}},
			history{Sequences: [][]int{
				{0, 3, 6, 9, 12, 15, 18},
				{3, 3, 3, 3, 3, 3},
				{0, 0, 0, 0, 0},
			}},
		},
		{
			history{Sequences: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			}},
			history{Sequences: [][]int{
				{1, 3, 6, 10, 15, 21, 28},
				{2, 3, 4, 5, 6, 7},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0},
			}},
		},
		{
			history{Sequences: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			}},
			history{Sequences: [][]int{
				{10, 13, 16, 21, 30, 45, 68},
				{3, 3, 5, 9, 15, 23},
				{0, 2, 4, 6, 8},
				{2, 2, 2, 2},
				{0, 0, 0},
			}},
		},
		{
			history{Sequences: [][]int{
				{-7, -12, -15, -4, 33},
				{-5, -3, 11, 37},
				{2, 14, 26},
				{12, 12},
				{0},
			}},
			history{Sequences: [][]int{
				{-7, -12, -15, -4, 33, 108},
				{-5, -3, 11, 37, 75},
				{2, 14, 26, 38},
				{12, 12, 12},
				{0, 0},
			}},
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			tc.in.extrapolate()
			got := tc.in
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
			input: strings.Split(test, "\n"),
			want:  114,
		},
		{
			input: strings.Split(input, "\n"),
			want:  1938731307,
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
			want:  2,
		},
		{
			input: strings.Split(input, "\n"),
			want:  948,
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
