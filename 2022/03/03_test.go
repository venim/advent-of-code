package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	input        string
	wantItem     string
	wantPriority int
}{
	{"vJrwpWtwJgWrhcsFMMfFFhFp", "p", 16},
	{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L", 38},
	{"PmmdzqPrVvPwwTWBwg", "P", 42},
	{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v", 22},
	{"ttgJtRGJQctTZtZT", "t", 20},
	{"CrZsJsPPZsGzwwsLwLmpwMDw", "s", 19},
}

func TestGetCommonItem(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			t.Logf("input len: %d", len(tc.input))
			got := getCommonItems(tc.input)
			want := []byte(tc.wantItem)
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("(%s) (-want, +got) \n%s", tc.input, diff)
			}
		})
	}
}

func TestScore(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := score([]byte(tc.wantItem))
			want := tc.wantPriority
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("(%s) (-want, +got) \n%s", tc.wantItem, diff)
			}
		})
	}
}

func TestSum(t *testing.T) {
	sacks := make([]string, 0, len(tests))
	got := 0
	want := 157
	for _, tc := range tests {
		sacks = append(sacks, tc.input)
	}
	for _, sack := range sacks {
		got += score(getCommonItems((sack)))
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

func TestGetBadge(t *testing.T) {
	badgeTests := []struct {
		elves []string
		want  string
	}{
		{[]string{tests[0].input, tests[1].input, tests[2].input}, "r"},
		{[]string{tests[3].input, tests[4].input, tests[5].input}, "Z"},
	}
	for _, tc := range badgeTests {
		t.Run("", func(t *testing.T) {
			got := getBadges(tc.elves...)
			want := []byte(tc.want)
			if diff := cmp.Diff(string(want), string(got)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
