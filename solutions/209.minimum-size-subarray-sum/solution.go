package solution

import "math"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	ans := math.MaxInt
	sum := 0
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
