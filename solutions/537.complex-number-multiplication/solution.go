package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	idx := strings.IndexByte(num1, '+')
	r1, _ := strconv.Atoi(num1[:idx])
	i1, _ := strconv.Atoi(num1[idx+1 : len(num1)-1])

	idx = strings.IndexByte(num2, '+')
	r2, _ := strconv.Atoi(num2[:idx])
	i2, _ := strconv.Atoi(num2[idx+1 : len(num2)-1])
	return fmt.Sprintf("%d+%di", r1*r2-i1*i2, r1*i2+i1*r2)
}
