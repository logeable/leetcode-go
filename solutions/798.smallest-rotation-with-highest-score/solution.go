package solution

func bestRotation(nums []int) int {
	var k int
	var max int
	for i := 0; i < len(nums); i++ {
		v := calc(rotation(nums, i))
		if v > max {
			max = v
			k = i
		}
	}
	return k
}

func rotation(nums []int, n int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = nums[(i+n)%len(nums)]
	}
	return ret
}

func calc(nums []int) int {
	var ret int
	for i, n := range nums {
		if n <= i {
			ret++
		}
	}
	return ret
}
