package solution

func missingRolls(rolls []int, mean int, n int) []int {
	total := (len(rolls) + n) * mean
	mtotal := 0
	for _, x := range rolls {
		mtotal += x
	}
	ntotal := total - mtotal
	min, max := n, 6*n

	if ntotal > max || ntotal < min {
		return nil
	}
	return nsum(ntotal, n)
}

func nsum(total int, n int) []int {
	x := total / n
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = x
	}

	r := (total - x*n)
	for r > 0 {
		for i := 0; i < n; i++ {
			ret[i]++
			r--
			if r == 0 {
				break
			}
		}
	}

	return ret
}
