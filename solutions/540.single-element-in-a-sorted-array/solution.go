package solution

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)>>1
		if m == l && m == r {
			return nums[m]
		}
		if (m-l)&1 == 1 {
			if nums[m] == nums[m-1] {
				l = m + 1
			} else if nums[m] == nums[m+1] {
				r = m - 1
			} else {
				return nums[m]
			}
		} else {
			if nums[m] == nums[m-1] {
				r = m - 2
			} else if nums[m] == nums[m+1] {
				l = m + 2
			} else {
				return nums[m]
			}
		}
	}
	panic("invalid input")
}
