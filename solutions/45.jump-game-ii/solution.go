package solution

func jump(nums []int) int {
	pos := len(nums) - 1
	var steps int
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				steps++
			}
		}
	}
	return steps
}
