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
			want:  46,
		},
		{
			input: strings.Split(input, "\n"),
			want:  6605,
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
			want:  51,
		},
		{
			input: strings.Split(input, "\n"),
			want:  6766,
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