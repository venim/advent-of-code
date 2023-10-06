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

func TestParseRates(t *testing.T) {
	lines := strings.Split(test, "\n")
	freq := parseFrequency(lines)
	gamma, epsilon := parseRates(freq)
	tests := []struct {
		gamma   int
		epsilon int
	}{
		{
			22,
			9,
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if diff := cmp.Diff(tc.gamma, gamma); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(tc.epsilon, epsilon); diff != "" {
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
		{
			input: strings.Split(test, "\n"),
			want:  198,
		},
		{
			input: strings.Split(input, "\n"),
			want:  3958484,
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

func TestFindRating(t *testing.T) {
	tests := []struct {
		name string
		want int
		fn   func([]string) int
	}{
		{"ogr", 23, findOxygenGeneratorRating},
		{"co2sr", 10, findCO2ScrubberRating},
	}
	lines := strings.Split(test, "\n")
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := tc.fn(lines)
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
			want:  230,
		},
		{
			input: strings.Split(input, "\n"),
			want:  1613181,
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
