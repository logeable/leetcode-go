package solution

func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if isPrime(bit1Count(i)) {
			count++
		}

	}
	return count
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func bit1Count(n int) int {
	ret := 0
	for n >= 1 {
		if n&1 == 1 {
			ret++
		}
		n >>= 1
	}
	return ret
}
