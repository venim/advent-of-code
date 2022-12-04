package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContainsSet(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"6-6", "4-6", true},
		{"4-6", "6-6", true},
		{"2-8", "3-7", true},
		{"2-4", "6-8", false},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := containsSet(tc.a, tc.b)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestSetsOverlap(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"2-3", "3-4", true},
		{"5-8", "1-6", true},
		{"2-8", "3-7", true},
		{"2-4", "6-8", false},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := setsOverlap(tc.a, tc.b)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
