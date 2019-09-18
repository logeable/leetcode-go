package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeOuterParentheses(t *testing.T) {
	table := []struct {
		input    string
		expected string
	}{
		{
			"(()())(())",
			"()()()",
		},
		{
			"(()())(())(()(()))",
			"()()()()(())",
		},
		{
			"()()",
			"",
		},
	}

	for _, row := range table {
		actual := removeOuterParentheses(row.input)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input: %q, actual: %q", row.input, actual))
	}
}
