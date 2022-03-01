package solution

import "math"

func concatenatedBinary(n int) int {
	ret := 0
	mod := (int(math.Pow10(9)) + 7)
	for i := 1; i <= n; i++ {
		ret = (ret << nbits(i)) + i
		ret %= mod
	}
	return ret
}
func nbits(n int) int {
	ret := 0
	for n > 0 {
		n /= 2
		ret++
	}
	return ret
}
