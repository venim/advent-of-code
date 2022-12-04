package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed in_test.txt
var testInput string

func TestMakSet(t *testing.T) {
	tests := []struct {
		sections string
		want     map[int64]bool
	}{
		{"2-4", map[int64]bool{2: true, 3: true, 4: true}},
		{"6-8", map[int64]bool{6: true, 7: true, 8: true}},
		{"2-3", map[int64]bool{2: true, 3: true}},
	}
	for _, tc := range tests {
		t.Run(tc.sections, func(t *testing.T) {
			got := makeSet(tc.sections)
			want := tc.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestContainsSet(t *testing.T) {
	tests := []struct {
		a, b map[int64]bool
		want bool
	}{
		{makeSet("6-6"), makeSet("4-6"), true},
		{makeSet("4-6"), makeSet("6-6"), true},
		{makeSet("2-8"), makeSet("3-7"), true},
		{makeSet("2-4"), makeSet("6-8"), false},
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
		a, b map[int64]bool
		want bool
	}{
		{makeSet("2-3"), makeSet("3-4"), true},
		{makeSet("5-8"), makeSet("1-6"), true},
		{makeSet("2-8"), makeSet("3-7"), true},
		{makeSet("2-4"), makeSet("6-8"), false},
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
