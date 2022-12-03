package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScoreGame(t *testing.T) {
	tests := []struct {
		guide     string
		wantScore int
	}{
		{"A Y", 8},
		{"B X", 1},
		{"C Z", 6},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := scoreGame(parsePart1(tc.guide))
			want := tc.wantScore
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestGuessGame(t *testing.T) {
	tests := []struct {
		guide     string
		wantScore int
	}{
		{"A Y", 4},
		{"B X", 1},
		{"C Z", 7},
	}
	for _, tc := range tests {
		t.Run(tc.guide, func(t *testing.T) {
			got := guessGame(parsePart2(tc.guide))
			want := tc.wantScore
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
