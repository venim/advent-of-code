package main

import (
	_ "embed"
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

func TestPart1(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"8A004A801A8002F478"},
			want:  16,
		},
		{
			input: []string{"620080001611562C8802118E34"},
			want:  12,
		},
		{
			input: []string{"C0015000016115A2E0802F182340"},
			want:  23,
		},
		{
			input: []string{"A0016C880162017C3686B18A3D4780"},
			want:  31,
		},
		{
			input: []string{input},
			want:  879,
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
			input: []string{"C200B40A82"},
			want:  3,
		},
		{
			input: []string{"04005AC33890"},
			want:  54,
		},
		{
			input: []string{"880086C3E88112"},
			want:  7,
		},
		{
			input: []string{"CE00C43D881120"},
			want:  9,
		},
		{
			input: []string{"D8005AC2A8F0"},
			want:  1,
		},
		{
			input: []string{"F600BC2D8F"},
			want:  0,
		},
		{
			input: []string{"9C005AC2F8F0"},
			want:  0,
		},
		{
			input: []string{"9C0141080250320F1802104A08"},
			want:  1,
		},
		{
			input: []string{input},
			want:  539051801941,
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
