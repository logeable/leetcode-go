package solution

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var ret []int
	for i, n := range nums {
		if n == target {
			ret = append(ret, i)
		}
	}
	return ret
}
