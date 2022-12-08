package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var test string = `30373
25512
65332
33549
35390`

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

func TestSurvey(t *testing.T) {
	tests := []struct {
		input     []string
		wantPart1 int
		wantPart2 int
	}{
		{strings.Split(test, "\n"), 21, 8},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			part1, part2 := survey(tc.input)
			if diff := cmp.Diff(tc.wantPart1, part1); diff != "" {
				t.Errorf("[Part 1]: %s", diff)
			}
			if diff := cmp.Diff(tc.wantPart2, part2); diff != "" {
				t.Errorf("[Part 2]: %s", diff)
			}
		})
	}
}
