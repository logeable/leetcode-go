package solution

func hasAlternatingBits(n int) bool {
	prev := n & 1
	n >>= 1
	for n > 0 {
		b := n & 1
		if b == prev {
			return false
		}
		prev = b
		n >>= 1
	}
	return true
}
