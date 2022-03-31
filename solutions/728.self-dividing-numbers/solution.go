package solution

func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for i := left; i <= right; i++ {
		ok := true
		for _, v := range allNum(i) {
			if v == 0 || i%v != 0 {
				ok = false
				break
			}
		}
		if ok {
			ret = append(ret, i)
		}
	}
	return ret
}
func allNum(v int) []int {
	var ret []int

	for v > 0 {
		ret = append(ret, v%10)
		v /= 10
	}
	return ret
}
