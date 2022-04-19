package solution

import (
	"sort"
	"strconv"
	"strings"
)

type LexInt []int

// Len is the number of elements in the collection.
func (l LexInt) Len() int {
	return len(l)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (l LexInt) Less(i int, j int) bool {
	return strings.Compare(strconv.Itoa(l[i]), strconv.Itoa(l[j])) < 0
}

// Swap swaps the elements with indexes i and j.
func (l LexInt) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func lexicalOrder(n int) []int {
	arr := make(LexInt, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	sort.Sort(&arr)
	return []int(arr)
}
