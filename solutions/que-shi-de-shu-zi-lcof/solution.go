package solution

func missingNumber(nums []int) int {
	for i, n := range nums {
		if i != n {
			return i
		}
	}
	return len(nums)
}
