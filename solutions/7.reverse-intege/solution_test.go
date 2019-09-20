package solution

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_overflow(t *testing.T) {
	t.Logf("max: %d\n", math.MaxInt32)
	t.Logf("min: %d\n", math.MinInt32)
}

func Test_reverse(t *testing.T) {
	table := []struct {
		input    int
		expected int
	}{
		{
			0,
			0,
		},
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
		{
			1534236469,
			0,
		},
	}
	for _, row := range table {
		actual := reverse(row.input)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input: %d, actual: %d", row.input, actual))
	}
}
