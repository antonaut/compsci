package main

import (
	"testing"
)

func TestStableMatch(t *testing.T) {
	t.Log("Det här är svinbra.")
	men_prefs := [][]int{
		[]int{},
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		[]int{1, 2, 3},
	}
	women_prefs := [][]int{
		[]int{},
		[]int{1, 2, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
	}
	n := 3

	matching := stable_match(men_prefs, women_prefs, n)

	for i := 1; i <= n; i++ {
		t.Logf("m: %d, w: %d", matching[i], i)
	}
}
