package solution

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l := findLeft(nums, target)
	if l == -1 {
		return 0
	}
	r := findRight(nums, target)
	if r == -1 {
		return 0
	}

	return r - l + 1
}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			if m == 0 || nums[m-1] != target {
				return m
			}
			r = m - 1
		}
	}
	return -1
}

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			if m == len(nums)-1 || nums[m+1] != target {
				return m
			}
			l = m + 1
		}
	}
	return -1
}
