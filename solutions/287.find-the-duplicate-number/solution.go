package solution

func findDuplicate(nums []int) int {
	s, f := 0, 0
	for s, f = nums[s], nums[nums[f]]; s != f; s, f = nums[s], nums[nums[f]] {
	}

	s = 0
	for s, f = nums[s], nums[f]; s != f; s, f = nums[s], nums[f] {
	}
	return s
}
