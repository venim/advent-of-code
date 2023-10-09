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

func TestParse(t *testing.T) {
	lines := strings.Split(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, "\n")
	want := caves{
		"start": {Id: "start", Connections: map[string]struct{}{"A": {}, "b": {}}},
		"A":     {Id: "A", Connections: map[string]struct{}{"start": {}, "b": {}, "c": {}, "end": {}}},
		"b":     {Id: "b", Connections: map[string]struct{}{"start": {}, "A": {}, "d": {}, "end": {}}},
		"c":     {Id: "c", Connections: map[string]struct{}{"A": {}}},
		"d":     {Id: "d", Connections: map[string]struct{}{"b": {}}},
		"end":   {Id: "end", Connections: map[string]struct{}{"A": {}, "b": {}}},
	}
	got := parse(lines)
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
			input: strings.Split(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, "\n"),
			want: 19,
		},
		{
			input: strings.Split(test, "\n"),
			want:  226,
		},
		{
			input: strings.Split(input, "\n"),
			want:  3485,
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
			input: strings.Split(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, "\n"),
			want: 36,
		},
		{
			input: strings.Split(test, "\n"),
			want:  3509,
		},
		{
			input: strings.Split(input, "\n"),
			want:  85062,
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
