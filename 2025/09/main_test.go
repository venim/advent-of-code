package main

import (
	_ "embed"
	"strings"
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

func TestCalc(t *testing.T) {
	tests := []struct {
		a    util.Pos
		b    util.Pos
		area int
	}{
		{util.Pos{2, 5}, util.Pos{9, 7}, 24},
		{util.Pos{9, 7}, util.Pos{2, 5}, 24},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := area(tc.a, tc.b)
			want := tc.area
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
			want:  50,
		},
		{
			input: strings.Split(input, "\n"),
			want:  4744899849,
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
			want:  24,
		},
		{
			input: strings.Split(input, "\n"),
			want:  0,
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
