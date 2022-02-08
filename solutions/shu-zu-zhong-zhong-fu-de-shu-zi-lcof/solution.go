package solution

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}

		m := nums[i]
		if nums[m] == m {
			return m
		} else {
			nums[i], nums[m] = nums[m], nums[i]
		}
	}
	return -1
}
