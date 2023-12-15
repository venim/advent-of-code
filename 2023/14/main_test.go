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

func TestRotate(t *testing.T) {
	tests := []struct {
		in   dish
		want dish
	}{
		{
			dish{{0, 1, 2}},
			dish{{0}, {1}, {2}},
		},
		{
			// 1 2 -> 3 1
			// 3 4    4 2
			dish{{1, 2}, {3, 4}},
			dish{{3, 1}, {4, 2}},
		},
		{
			// 1 2 3 -> 4 1
			// 4 5 6    5 2
			//          6 3
			dish{{1, 2, 3}, {4, 5, 6}},
			dish{{4, 1}, {5, 2}, {6, 3}},
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := tc.in
			got.rotate()
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
			want:  136,
		},
		{
			input: strings.Split(input, "\n"),
			want:  105249,
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
			want:  64,
		},
		{
			input: strings.Split(input, "\n"),
			want:  88680,
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
