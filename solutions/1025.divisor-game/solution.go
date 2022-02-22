package solution

func divisorGame(n int) bool {
	f := make([]bool, n+2)
	f[1], f[2] = false, true

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
