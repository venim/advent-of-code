package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	//go:embed in_test.txt
	test        string
	conversions = []struct {
		snafu string
		dec   int
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}
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

func BenchmarkStoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stoi(conversions[0].snafu)
	}
}

func TestStoi(t *testing.T) {
	for _, tc := range conversions {
		t.Run("", func(t *testing.T) {
			got := stoi(tc.snafu)
			want := tc.dec
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestItos(t *testing.T) {
	for _, tc := range conversions {
		t.Run("", func(t *testing.T) {
			got := itos(tc.dec)
			want := tc.snafu
			if diff := cmp.Diff(want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{strings.Split(test, "\n"), "2=-1=0"},
		{strings.Split(input, "\n"), "2-00=12=21-0=01--000"},
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
