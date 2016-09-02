package main

type node struct {
	man  int
	next *node
}

// This is a `\Theta(n^2)` implementation of a solution to the stable matching
// (marriage) problem. More info at:
// https://en.wikipedia.org/wiki/Stable_marriage_problem

// Women and men are represented as int\`s from `[1..n]` both inclusive. Every
// woman needs to rank every man and vice versa in their preference lists. The
// preferences are ordered such that the most wanted person comes first in the
// list.

// This function takes preference lists from both men and women and returns a
// slice `husband[w]` saying that woman `w` should be married to husband
// `husband[w]`.
func stable_match(men_pref [][]int, women_pref [][]int) []int {

	n := len(men_pref)

	next := make([]int, n+1)
	current := make([]int, n+1)
	rank := make([][]int, n+1)

	var free_men *node

	for i := 1; i <= n; i++ {
		next[i] = 1
		current[i] = 0
		rank[i] = make([]int, n+1)
		man_node := &node{i, free_men}
		free_men = man_node

		for j := 0; j < n; j++ {
			rank[i][women_pref[i-1][j]] = j
		}
	}

	for free_men != nil {
		man := free_men.man
		target_bride := men_pref[man-1][next[man]]
		next[man]++
		if current[target_bride] == 0 {
			current[target_bride] = man
			free_men = free_men.next
			continue
		}

		if rank[target_bride][man] < rank[target_bride][current[target_bride]] {
			tmp := current[target_bride]
			current[target_bride] = man
			man_node := &node{tmp, free_men}
			free_men = man_node
		}
	}

	return current
}

func main() {

}
