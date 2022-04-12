package solution

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	c := nums[l]
	ans := 10000
	count := 1
	found := false
	for l <= r && r < len(nums) {
		if c < target {
			r++
			if r == len(nums) {
				break
			}
			c += nums[r]
		} else {
			found = true
			count = r - l + 1
			if count < ans {
				ans = count
			}
			c -= nums[l]
			l++
		}
	}
	if !found {
		return 0
	}
	return ans
}
