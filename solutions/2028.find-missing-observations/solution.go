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

	v, r := ntotal/n, ntotal%n
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = v
		if i < r {
			ret[i]++
		}
	}
	return ret
}
