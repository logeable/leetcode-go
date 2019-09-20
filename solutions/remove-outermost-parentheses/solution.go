package solution

import "bytes"

func removeOuterParentheses(S string) string {
	if S == "" {
		return ""
	}

	stack := make([]rune, 0, len(S))
	var primitives []string

	var last int
	for i, c := range S {
		if c == '(' {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			primitives = append(primitives, S[last:i+1])
			last = i + 1
		}
	}
	var ans bytes.Buffer
	for _, x := range primitives {
		if x == "()" {
			continue
		}
		ans.WriteString(x[1 : len(x)-1])
	}

	return ans.String()
}
