package sorting

import "testing"

func TestQuickSort(t *testing.T) {
	sq := *new(QuickSorter)
	A := []int{2, 4, 3, 5, 7, 2, 2, 1, -2, 9}
	sq.Sort(A, func(a int, b int) bool { return a < b })
}

func TestSorting(t *testing.T) {
	var A []int
	sorters := [...]Sorter{
		*new(BubbleSorter),
		*new(InsertionSorter),
		*new(SelectionSorter),
		*new(MergeSorter),
		*new(QuickSorter),
	}

	for _, s := range sorters {
		A = []int{2, 4, 3, 5, 7, 2, 2, 1, -2, 9}
		s.Sort(A, func(a int, b int) bool {
			return a < b
		})
		for i, e := range []int{-2, 1, 2, 2, 2, 3, 4, 5, 7, 9} {
			if A[i] != e {
				t.Errorf("A[i]: %v !=  e: %v\n", A[i], e)
			}
		}
		t.Logf("%v\n", A)
	}
}
