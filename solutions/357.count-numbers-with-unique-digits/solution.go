package solution

func countNumbersWithUniqueDigits(n int) int {
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	ans := 0
	if max > 0 {
		ans++
	}
	for i := 1; i < max; i++ {
		cache := make(map[int]struct{})
		n := i
		found := false
		for n > 0 {
			m := n % 10
			n /= 10
			if _, ok := cache[m]; ok {
				found = true
				break
			}
			cache[m] = struct{}{}
		}
		if !found {
			ans++
		}
	}

	return ans
}
