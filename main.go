package main

import "./sorting"

func main() {
	sq := *new(sorting.QuickSorter)
	A := []int{2, 3, 1, 5, 6}
	sq.Sort(A, func(a, b int) bool { return a < b })
}
