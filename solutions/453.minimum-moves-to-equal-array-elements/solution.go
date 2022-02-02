package solution

func minMoves(nums []int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	ans := 0
	for _, n := range nums {
		ans += (n - min)
	}
	return ans
}
