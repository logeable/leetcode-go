package solution

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i, x := range nums {
		if v, ok := cache[target-x]; ok {
			return []int{v, i}
		}
		cache[x] = i
	}
	return nil
}
