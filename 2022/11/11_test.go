package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed in_test.txt
var test string

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

func TestParseItems(t *testing.T) {
	tests := []struct {
		line string
		want []int
	}{
		{"Starting items: 79, 98", []int{79, 98}},
		{"  Starting items: 54, 65, 75, 74", []int{54, 65, 75, 74}},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := parseItems(tc.line)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestMakeOp(t *testing.T) {
	tests := []struct {
		line string
		old  int
		want int
	}{
		{"  Operation: new = old * 19", 1, 19},
		{"  Operation: new = old * 19", 2, 19 * 2},
		{"  Operation: new = old * old", 2, 4},
		{"  Operation: new = old + 20", 2, 22},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := makeOp(tc.line)(tc.old)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input  []string
		rounds int
		part   int
		want   int
	}{
		{strings.Split(test, "\n"), 20, 1, 10605},
		{strings.Split(test, "\n"), 10000, 2, 2713310158},
		{strings.Split(input, "\n"), 20, 1, 117624},
		{strings.Split(input, "\n"), 10000, 2, 16792940265},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := solve(tc.input, tc.rounds, tc.part)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("[Part %d]: \n%s", tc.part, diff)
			}
		})
	}
}
