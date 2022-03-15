package solution

func getMaxLen(nums []int) int {
	positive := make([]int, len(nums))
	negative := make([]int, len(nums))

	if nums[0] > 0 {
		positive[0] = 1
	} else if nums[0] < 0 {
		negative[0] = 1
	}
	maxLen := positive[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positive[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				negative[i] = negative[i-1] + 1
			} else {
				negative[i] = 0
			}
		} else if nums[i] < 0 {
			negative[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				positive[i] = negative[i-1] + 1
			} else {
				positive[i] = 0
			}
		} else {
			positive[i] = 0
			negative[i] = 0
		}

		if positive[i] > maxLen {
			maxLen = positive[i]
		}
	}

	return maxLen
}
