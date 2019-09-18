package solution

func reverse(x int) int {

	var ret int

	for x != 0 {
		m := x % 10
		x = x / 10
		ret = ret*10 + m
	}

	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}
