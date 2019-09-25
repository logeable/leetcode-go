package solution

import (
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	var i int
	var tmp string
	l := (len(B) + len(A)) / len(A)
	for i <= l {
		tmp += A
		i++

		if strings.Index(tmp, B) != -1 {
			return i
		}
	}
	return -1
}
