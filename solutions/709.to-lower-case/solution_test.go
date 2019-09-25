package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toLowerCase(t *testing.T) {
	table := []struct {
		input    string
		expected string
	}{
		{
			"Hello",
			"hello",
		},
		{
			"here",
			"here",
		},
		{
			"",
			"",
		},
		{
			"LOVELY",
			"lovely",
		},
	}

	for _, row := range table {
		assert.Equal(t, row.expected, toLowerCase(row.input), fmt.Sprintf("input: %q", row.input))
	}
}
