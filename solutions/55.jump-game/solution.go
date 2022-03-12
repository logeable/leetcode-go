package solution

func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		if v := i + nums[i]; v > k {
			k = v
		}
		if k >= len(nums)-1 {
			return true
		}
	}
	return true
}
