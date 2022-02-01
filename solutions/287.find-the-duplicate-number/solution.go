package solution

import "sort"

func findDuplicate(nums []int) int {

	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)
	for i := 0; i < len(newNums)-1; i++ {
		if newNums[i] == newNums[i+1] {
			return newNums[i]
		}
	}
	panic("invalid test")
}
