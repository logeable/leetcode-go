package solution

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repeatedStringMatch(t *testing.T) {
	table := []struct {
		A        string
		B        string
		expected int
	}{
		{
			"abcd",
			"cdabcdab",
			3,
		},
		{
			"abcd",
			"cdabc",
			2,
		},
		{
			"abcd",
			"abcd",
			1,
		},
		{
			"abcd",
			"abc",
			1,
		},
		{
			"abcd",
			"dbc",
			-1,
		},
		{
			"a",
			strings.Repeat("a", 1000),
			1000,
		},
	}

	for _, row := range table {
		actual := repeatedStringMatch(row.A, row.B)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("A: %q, B: %q", row.A, row.B))
	}
}
