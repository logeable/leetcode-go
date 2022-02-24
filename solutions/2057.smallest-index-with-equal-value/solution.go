package solution

func smallestEqual(nums []int) int {
	ret := -1
	for i, n := range nums {
		if i%10 == n {
			return i
		}
	}
	return ret
}
