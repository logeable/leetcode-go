package solution

func maxSubarraySumCircular(nums []int) int {
	total := nums[0]

	maxDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	maxSum := nums[0]

	minDp := make([]int, len(nums))
	minDp[0] = nums[0]
	minSum := nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		maxDp[i] = max(nums[i], maxDp[i-1]+nums[i])
		maxSum = max(maxDp[i], maxSum)
		minDp[i] = min(nums[i], minDp[i-1]+nums[i])
		minSum = min(minDp[i], minSum)
	}

	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, total-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
