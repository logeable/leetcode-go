package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	table := []struct {
		input    []interface{}
		expected []int
	}{
		{
			[]interface{}{1, nil, 2, 3},
			[]int{1, 3, 2},
		},
	}

	for _, row := range table {
		root := types.MakeTreeNode(row.input)
		actual := inorderTraversal(root)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input: %v, actual: %v", row.input, actual))
	}
}
