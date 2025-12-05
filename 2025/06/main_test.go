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

func TestParseNumbers(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		{" 1 2 3", []int{1, 2, 3}},
		{"  6 98  215 314", []int{6, 98, 215, 314}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := parseNumbers(tc.in)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestParsOperands(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{" *   +   *   + ", []string{"*", "+", "*", "+"}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := parseOperands(tc.in)
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
			want:  4277556,
		},
		{
			input: strings.Split(input, "\n"),
			want:  4412382293768,
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

func TestParseColumns(t *testing.T) {
	tests := []struct {
		in   []string
		want []int
	}{
		{[]string{"64 ", "23 ", "314"}, []int{4, 431, 623}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := parseColumns(tc.in)
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
			want:  3263827,
		},
		{
			input: strings.Split(input, "\n"),
			want:  7858808482092,
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
