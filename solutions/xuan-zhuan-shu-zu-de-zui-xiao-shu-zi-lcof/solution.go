package solution

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		m := l + (r-l)/2
		if numbers[m] < numbers[r] {
			r = m
		} else if numbers[m] > numbers[r] {
			l = m + 1
		} else {
			r--
		}
	}
	return numbers[r]
}
