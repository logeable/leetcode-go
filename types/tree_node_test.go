package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeNode_String(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Left = &TreeNode{4, nil, nil}
	t.Log("\n" + root.String())
}

func TestMakeTreeNode(t *testing.T) {
	table := []struct {
		input []interface{}
	}{
		{
			[]interface{}{1, nil, 2, 3},
		},
		{
			[]interface{}{1, 2, 3, 4},
		},
		{
			[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			[]interface{}{1, nil, 2, 3, nil, nil, 4, 5, 6},
		},
	}
	for _, row := range table {
		root := MakeTreeNode(row.input)
		t.Log("\n" + root.String())
	}
}

func TestMakeTreeNode_invalidData(t *testing.T) {
	input := []interface{}{1, nil, nil, 3}
	assert.Panics(t, func() {
		MakeTreeNode(input)
	}, fmt.Sprintf("input: %v", input))
}
