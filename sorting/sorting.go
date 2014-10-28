package sorting

import (
	"math"
)

type Sorter interface {
	Sort([]int, func(int, int) bool)
}

type BubbleSorter struct{}

// Don't use this. Ever.
func (BubbleSorter) Sort(slice []int, cmp func(int, int) bool) {
	//fmt.Printf("slice: %v, len: %v\n", slice, len(slice))
	for {
		notSwapped := true
		for i := len(slice) - 2; i >= 0; i-- {
			if cmp(slice[i+1], slice[i]) {
				//fmt.Printf("  a: %v, b: %v\n", slice[i+1], slice[i])
				slice[i], slice[i+1] = slice[i+1], slice[i]
				notSwapped = false
			}
		}
		if notSwapped {
			break
		}
	}
}

type InsertionSorter struct{}

// Worst case: T(n) => O(n^2)
// Best case: T(n) => O(n)
func (InsertionSorter) Sort(slice []int, cmp func(int, int) bool) {
	end := len(slice)
	for end > 0 {
		extremeIndex := 0
		for j := 0; j < end; j++ {
			if cmp(slice[extremeIndex], slice[j]) {
				extremeIndex = j
			}
		}
		slice[extremeIndex], slice[end-1] = slice[end-1], slice[extremeIndex]
		end--
	}
}

type SelectionSorter struct{}

// T(n) => O(n^2)
func (SelectionSorter) Sort(slice []int, cmp func(int, int) bool) {
	for i := 0; i < len(slice)-1; i++ {
		// invariant; slice[x] | x < i is sorted
		minidx := i
		for j := i + 1; j < len(slice); j++ {
			if cmp(slice[j], slice[minidx]) {
				minidx = j
			}
		}
		slice[i], slice[minidx] = slice[minidx], slice[i]
	}
}

type MergeSorter struct {
	cmp   func(int, int) bool
	slice []int
}

// Now we're starting to get somewhere.
// Worst case: T(n) => O(n*log(n))
func (ms MergeSorter) Sort(slice []int, cmp func(int, int) bool) {
	ms.cmp = cmp
	ms.slice = slice
	ms.mergeSort(0, len(slice)-1)
}

func (ms MergeSorter) mergeSort(p, r int) {
	if p < r {
		q := (p + r) / 2
		ms.mergeSort(p, q)
		ms.mergeSort(q+1, r)
		ms.merge(p, q, r)
	}
}

func (ms MergeSorter) merge(p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < len(L)-1; i++ {
		L[i] = ms.slice[p+i]
	}

	for j := 0; j < len(R)-1; j++ {
		R[j] = ms.slice[q+j+1]
	}

	L[n1] = math.MaxInt64
	R[n2] = math.MaxInt64
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if ms.cmp(L[i], R[j]) {
			ms.slice[k] = L[i]
			i++
		} else {
			ms.slice[k] = R[j]
			j++
		}
	}
}

type QuickSorter struct{}

// TODO
func (qs QuickSorter) Sort(slice []int, cmp func(int, int) bool) {
	// pick pivot, median of three

	// partition around pivot, divide and conquer until insertion sort at fixed LIMIT
}

type RadixSorter struct{}

// TODO
func (rs RadixSorter) Sort(slice []int, cmp func(int, int) bool) {

}
