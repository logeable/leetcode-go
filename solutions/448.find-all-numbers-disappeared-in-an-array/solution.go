package solution

func findDisappearedNumbers(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); {
		if nums[i] == nums[nums[i]-1] {
			i++
		} else {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, x := range nums {
		if x != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
