package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	nums   []int
	target int
}

var (
	table = []struct {
		param    param
		expected []int
	}{
		{
			param{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{0, 1},
		},
		{
			param{
				[]int{1, 5, 9, 20},
				29,
			},
			[]int{2, 3},
		},
		{
			param{
				[]int{1, 5, 9, 20},
				2,
			},
			nil,
		},
	}
)

func Test_twoSumBruteForce(t *testing.T) {

	for _, row := range table {
		assert.Equal(t, row.expected, twoSumBruteForce(row.param.nums, row.param.target))
	}
}

func Benchmark_twoSumBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumBruteForce(table[0].param.nums, table[0].param.target)
	}
}

func Test_twoSum(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, twoSum(row.param.nums, row.param.target))
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum(table[0].param.nums, table[0].param.target)
	}
}
