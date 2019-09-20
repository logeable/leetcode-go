package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	table = []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
	}
)

func Test_findMedianSortedArrays(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, findMedianSortedArrays(row.nums1, row.nums2), fmt.Sprintf("input: %v %v", row.nums1, row.nums2))
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(table[0].nums1, table[0].nums2)
	}
}
