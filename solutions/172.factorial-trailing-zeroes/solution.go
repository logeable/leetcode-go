package solution

func trailingZeroes(n int) int {
	var count int
	for n > 0 {
		count += n / 5
		n /= 5
	}

	return count
}
