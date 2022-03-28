package solution

func hasAlternatingBits(n int) bool {
	bits := tobits(n)
	if len(bits) == 1 {
		return true
	}

	for i := 1; i < len(bits); i++ {
		if bits[i] == bits[i-1] {
			return false
		}
	}
	return true
}

func tobits(n int) []byte {
	var ret []byte
	for n > 0 {
		ret = append(ret, byte(n&1))
		n >>= 1
	}
	return ret
}
