package main

type node struct {
	man  int
	next *node
}

// This is an Theta(n^2) implementation of a solution to the stable matching
// problem
func stable_match(men_pref [][]int, women_pref [][]int, n int) []int {
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
			rank[i][women_pref[i][j]] = j
		}
	}

	for free_men != nil {
		man := free_men.man
		target_bride := men_pref[man][next[man]]
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
