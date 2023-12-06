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

func TestAmConvert(t *testing.T) {
	tests := []struct {
		am   Map
		in   int
		want int
	}{
		{
			Map{rules: []Rule{
				{50, 98, 2},
				{52, 50, 48},
			}}, 50, 52,
		},
		{
			Map{rules: []Rule{
				{50, 98, 2},
				{52, 50, 48},
			}}, 49, 49,
		},
		{
			Map{rules: []Rule{
				{50, 98, 2},
				{52, 50, 48},
			}}, 98, 50,
		},
		{
			Map{rules: []Rule{
				{50, 98, 2},
				{52, 50, 48},
			}}, 97, 99,
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := tc.am.convert(tc.in)
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
			want:  35,
		},
		{
			input: strings.Split(input, "\n"),
			want:  579439039,
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
			want:  46,
		},
		{
			input: strings.Split(input, "\n"),
			want:  7873084,
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
