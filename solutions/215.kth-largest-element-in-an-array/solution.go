package solution

import "math/rand"

func findKthLargest(nums []int, k int) int {
	rk := len(nums) - k
	start, end := 0, len(nums)
	for {
		m := partition(nums, start, end)
		if m < rk {
			start = m + 1
		} else if m > rk {
			end = m
		} else {
			break
		}
	}
	return nums[rk]
}

func partition(nums []int, start, end int) int {
	nums = nums[start:end]
	i := rand.Intn(len(nums))
	v := nums[i]
	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]

	small := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < v {
			small++
			nums[small], nums[i] = nums[i], nums[small]
		}
	}
	small++
	nums[small], nums[len(nums)-1] = nums[len(nums)-1], nums[small]
	return small + start
}
