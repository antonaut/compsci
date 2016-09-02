package main

import (
	"testing"
)

func TestStableMatch(t *testing.T) {
	men_prefs := [][]int{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		[]int{1, 2, 3},
	}
	women_prefs := [][]int{
		[]int{1, 2, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
	}
	n := len(men_prefs)

	matching := stable_match(men_prefs, women_prefs)

	for i := 1; i <= n; i++ {
		t.Logf("m: %d, w: %d", matching[i], i)
	}
}
