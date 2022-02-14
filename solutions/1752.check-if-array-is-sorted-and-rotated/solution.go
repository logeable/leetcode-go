package solution

func check(nums []int) bool {
	var i = 1
	for i = 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			break
		}
	}
	if i == len(nums) {
		return true
	}

	for i = i + 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return nums[len(nums)-1] <= nums[0]
}
