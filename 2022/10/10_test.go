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

func TestCRT(t *testing.T) {
	tests := []struct {
		input     []string
		cycles    []int
		wantPart1 int
		wantPart2 string
	}{
		{strings.Split(test, "\n"), []int{20}, 420, "##..##..##..##..##.."},
		{strings.Split(test, "\n"), []int{21}, 441, "##..##..##..##..##..#"},
		{strings.Split(test, "\n"), []int{20, 60, 100, 140, 180, 220}, 13140, ""},
		{strings.Split(input, "\n"), []int{20, 60, 100, 140, 180, 220}, 13680, ""},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			part1, part2 := run(tc.input, tc.cycles)
			if diff := cmp.Diff(tc.wantPart1, part1); diff != "" {
				t.Errorf("[Part 1]: %s", diff)
			}
			n := tc.cycles[len(tc.cycles)-1]
			if diff := cmp.Diff(tc.wantPart2, part2[:n]); tc.wantPart2 != "" && diff != "" {
				t.Errorf("[Part 2]: %s", diff)
			}
			t.Logf("\n%s", part2)
		})
	}
}
