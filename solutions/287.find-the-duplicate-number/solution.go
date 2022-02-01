package solution

func findDuplicate(nums []int) int {
	for i := 1; i < len(nums); {
		m := nums[i]
		if m != i {
			if nums[m] != m {
				nums[m], nums[i] = nums[i], nums[m]
			} else {
				return m
			}
		} else {
			i++
		}
	}
	return nums[0]
}
