package solution

import "fmt"

func findKthNumber(n int, k int) int {
	prefix := 1
	p := 1
	for p < k {
		count := getNum(prefix, n)
		if k-p <= count {
			prefix *= 10
			p++
		} else {
			prefix += 1
			p += (count + 1)
		}
	}
	return prefix
}

func getNum(prefix int, n int) int {
	cur := prefix * 10
	next := (prefix + 1) * 10
	count := 0
	for cur <= n {
		count += min(next, n+1) - cur
		cur *= 10
		next *= 10
	}
	fmt.Println(prefix, count)
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
