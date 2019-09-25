package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	table := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
	}

	for _, row := range table {
		assert.Equal(t, row.expected, findDisappearedNumbers(row.input), fmt.Sprintf("input: %v", row.input))
	}
}
