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

func TestSolve(t *testing.T) {
	tests := []struct {
		input     []string
		row       int
		wantPart1 int
		tuning    int
		wantPart2 int
	}{
		{strings.Split(test, "\n"), 10, 26, 20, 56000011},
		{strings.Split(input, "\n"), 2000000, 5838453, 4000000, 12413999391794},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			part1, part2 := solve(tc.input, tc.row, tc.tuning)
			if diff := cmp.Diff(tc.wantPart1, part1); diff != "" {
				t.Errorf("[Part 1]: \n%s", diff)
			}
			if diff := cmp.Diff(tc.wantPart2, part2); diff != "" {
				t.Errorf("[Part 2]: \n%s", diff)
			}
		})
	}
}
