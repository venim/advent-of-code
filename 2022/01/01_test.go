package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed in_test.txt
var testInput string

func TestParseInventory(t *testing.T) {
	want := []int64{6000, 4000, 11000, 24000, 10000}
	got := parseInventory(testInput)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}
