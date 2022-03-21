package solution

func winnerOfGame(colors string) bool {
	cur := colors[0]
	cnt := 1
	var freq [2]int

	for i := 1; i < len(colors); i++ {
		c := colors[i]
		if c == cur {
			cnt++
			if cnt >= 3 {
				freq[cur-'A']++
			}
		} else {
			cnt = 1
			cur = c
		}
	}
	return freq[0] > freq[1]
}
