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

func TestBlizzardMove(t *testing.T) {
	tests := []struct {
		desc     string
		blizzard *blizzard
		gridSize pos
		wantPos  pos
	}{
		{
			">",
			&blizzard{pos{0, 0}, pos{0, 1}, '>'},
			pos{4, 6},
			pos{0, 1},
		},
		{
			"> wrap",
			&blizzard{pos{0, 2}, pos{0, 1}, '>'},
			pos{2, 3},
			pos{0, 0},
		},
		{
			"<",
			&blizzard{pos{0, 3}, pos{0, -1}, '<'},
			pos{4, 6},
			pos{0, 2},
		},
		{
			"< wrap",
			&blizzard{pos{0, 0}, pos{0, -1}, '<'},
			pos{1, 3},
			pos{0, 2},
		},
		{
			"^",
			&blizzard{pos{1, 4}, pos{-1, 0}, '^'},
			pos{4, 6},
			pos{0, 4},
		},
		{
			"^ wrap",
			&blizzard{pos{0, 4}, pos{-1, 0}, '^'},
			pos{5, 5},
			pos{4, 4},
		},
		{
			"v",
			&blizzard{pos{0, 4}, pos{1, 0}, 'v'},
			pos{4, 6},
			pos{1, 4},
		},
		{
			"v wrap",
			&blizzard{pos{2, 1}, pos{1, 0}, 'v'},
			pos{3, 2},
			pos{0, 1},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			b := tc.blizzard
			b.move(tc.gridSize)
			got := b.Pos
			want := tc.wantPos
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{strings.Split(test, "\n"), 18},
		// TODO: figure out why this fails...
		{strings.Split(input, "\n"), 238},
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
		{strings.Split(test, "\n"), 54},
		{strings.Split(input, "\n"), 751},
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
