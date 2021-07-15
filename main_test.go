package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	out := maxAreaOfIsland(in)
	want := 6
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}

	in = [][]int{{1, 1}}
	out = maxAreaOfIsland(in)
	want = 2
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
