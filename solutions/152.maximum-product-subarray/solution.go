package solution

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	maxVal := dpMax[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max3(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i], nums[i])
		dpMin[i] = min3(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i], nums[i])
		if dpMax[i] > maxVal {
			maxVal = dpMax[i]
		}
	}
	return maxVal
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}
func min3(a, b, c int) int {
	return min(min(a, b), c)
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
