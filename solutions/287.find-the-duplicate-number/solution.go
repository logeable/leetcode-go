package solution

func findDuplicate(nums []int) int {
	cache := make(map[int]int)
	for _, n := range nums {
		cache[n]++
		if cache[n] > 1 {
			return n
		}
	}
	panic("invalid nums")
}
