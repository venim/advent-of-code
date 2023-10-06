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

func TestParseInput(t *testing.T) {
	want := &data{
		Numbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		Boards: []board{
			{
				Spots: map[pos]bool{},
				Numbers: map[int]pos{
					22: {0, 0}, 13: {1, 0}, 17: {2, 0}, 11: {3, 0}, 0: {4, 0},
					8: {0, 1}, 2: {1, 1}, 23: {2, 1}, 4: {3, 1}, 24: {4, 1},
					21: {0, 2}, 9: {1, 2}, 14: {2, 2}, 16: {3, 2}, 7: {4, 2},
					6: {0, 3}, 10: {1, 3}, 3: {2, 3}, 18: {3, 3}, 5: {4, 3},
					1: {0, 4}, 12: {1, 4}, 20: {2, 4}, 15: {3, 4}, 19: {4, 4},
				},
			},
			{
				Spots: map[pos]bool{},
				Numbers: map[int]pos{
					3: {0, 0}, 15: {1, 0}, 0: {2, 0}, 2: {3, 0}, 22: {4, 0},
					9: {0, 1}, 18: {1, 1}, 13: {2, 1}, 17: {3, 1}, 5: {4, 1},
					19: {0, 2}, 8: {1, 2}, 7: {2, 2}, 25: {3, 2}, 23: {4, 2},
					20: {0, 3}, 11: {1, 3}, 10: {2, 3}, 24: {3, 3}, 4: {4, 3},
					14: {0, 4}, 21: {1, 4}, 16: {2, 4}, 12: {3, 4}, 6: {4, 4},
				},
			},
			{
				Spots: map[pos]bool{},
				Numbers: map[int]pos{
					14: {0, 0}, 21: {1, 0}, 17: {2, 0}, 24: {3, 0}, 4: {4, 0},
					10: {0, 1}, 16: {1, 1}, 15: {2, 1}, 9: {3, 1}, 19: {4, 1},
					18: {0, 2}, 8: {1, 2}, 23: {2, 2}, 26: {3, 2}, 20: {4, 2},
					22: {0, 3}, 11: {1, 3}, 13: {2, 3}, 6: {3, 3}, 5: {4, 3},
					2: {0, 4}, 0: {1, 4}, 12: {2, 4}, 3: {3, 4}, 7: {4, 4},
				},
			},
		},
	}

	got := parseInput(strings.Split(test, "\n"))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}

func TestWon(t *testing.T) {
	b := board{
		Spots: map[pos]bool{},
		Numbers: map[int]pos{
			14: {0, 0}, 21: {1, 0}, 17: {2, 0}, 24: {3, 0}, 4: {4, 0},
			10: {0, 1}, 16: {1, 1}, 15: {2, 1}, 9: {3, 1}, 19: {4, 1},
			18: {0, 2}, 8: {1, 2}, 23: {2, 2}, 26: {3, 2}, 20: {4, 2},
			22: {0, 3}, 11: {1, 3}, 13: {2, 3}, 6: {3, 3}, 5: {4, 3},
			2: {0, 4}, 0: {1, 4}, 12: {2, 4}, 3: {3, 4}, 7: {4, 4},
		},
	}
	won := false
	for _, i := range []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24} {
		if won == true {
			t.Error("board should not have won yet")
		}
		won = b.callN(i)
	}
	if won != true {
		t.Error("board should have won")
	}
}

func TestSum(t *testing.T) {
	b := board{
		Spots: map[pos]bool{},
		Numbers: map[int]pos{
			14: {0, 0}, 21: {1, 0}, 17: {2, 0}, 24: {3, 0}, 4: {4, 0},
			10: {0, 1}, 16: {1, 1}, 15: {2, 1}, 9: {3, 1}, 19: {4, 1},
			18: {0, 2}, 8: {1, 2}, 23: {2, 2}, 26: {3, 2}, 20: {4, 2},
			22: {0, 3}, 11: {1, 3}, 13: {2, 3}, 6: {3, 3}, 5: {4, 3},
			2: {0, 4}, 0: {1, 4}, 12: {2, 4}, 3: {3, 4}, 7: {4, 4},
		},
	}
	for _, i := range []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24} {
		b.callN(i)
	}
	want := 188
	got := b.sum()
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
			input: strings.Split(test, "\n"),
			want:  4512,
		},
		{
			input: strings.Split(input, "\n"),
			want:  38594,
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
			want:  1924,
		},
		{
			input: strings.Split(input, "\n"),
			want:  21184,
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
