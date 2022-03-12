package solution

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	for i, n := range nums {
		if i == len(nums)-1 {
			return true
		}
		if n == 0 {
			found := false
			for j := 0; j < i; j++ {
				if nums[j] > i-j {
					found = true
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
