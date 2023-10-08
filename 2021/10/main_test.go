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

func TestCorrupted(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"(]"},
		{"{()()()>"},
		{"(((()))}"},
		{"<([]){()}[{}])"},
		{"{([(<{}[<>[]}>{[]{[(<()>"},
		{"[[<[([]))<([[{}[[()]]]"},
		{"[{[{({}]{}}([{[{{{}}([]"},
		{"[<(<(<(<{}))><([]([]()"},
		{"<{([([[(<>()){}]>(<<{{"},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if _, ok := parseLine(tc.in); !ok {
				t.Errorf("%v should be corrupt", tc.in)
			}
		})
	}
}
func TestNotCorrupted(t *testing.T) {
	tests := []struct {
		in string
	}{
		{"([])"},
		{"{()()()}"},
		{"(((())))"},
		{"<([]){()}[{}]>"},
		{"(((((((((())))))))))"},
		{"[<>({}){}[([])<>]]"},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if _, ok := parseLine(tc.in); ok {
				t.Errorf("%v should not be corrupt", tc.in)
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
			want:  26397,
		},
		{
			input: strings.Split(input, "\n"),
			want:  343863,
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

func TestComplete(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "[({(<(())[]>[[{[]{<()<>>", want: "}}]])})]"},
		{in: "[(()[<>])]({[<{<<[]>>(", want: ")}>]})"},
		{in: "(((({<>}<{<{<>}{[]{[]{}", want: "}}>}>))))"},
		{in: "{<[[]]>}<{[{[{[]{()[[[]", want: "]]}}]}]}>"},
		{in: "<{([{{}}[<[[[<>{}]]]>[]]", want: "])}>"},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			stack, _ := parseLine(tc.in)
			got := complete([]rune(stack))
			if diff := cmp.Diff(tc.want, string(got)); diff != "" {
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
			want:  288957,
		},
		{
			input: strings.Split(input, "\n"),
			want:  2924734236,
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
