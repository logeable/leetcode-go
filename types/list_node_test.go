package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNode_String(t *testing.T) {
	table := []struct {
		node     *ListNode
		expected string
	}{
		{
			nil,
			"()",
		},
		{
			&ListNode{
				1,
				nil,
			},
			"(1)",
		},
		{
			&ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
			"(1 -> 2)",
		},
	}

	for _, row := range table {
		assert.Equal(t, row.expected, row.node.String(), fmt.Sprintf("input: %v", row.node))
	}
}

func TestMakeListNode(t *testing.T) {
	table := []struct {
		arg      []int
		expected *ListNode
	}{
		{
			nil,
			nil,
		},
		{
			[]int{},
			nil,
		},
		{
			[]int{1},
			&ListNode{
				1,
				nil,
			},
		},
		{
			[]int{1, 2, 3},
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
		},
	}
	for _, row := range table {
		assert.Equal(t, row.expected, MakeListNode(row.arg), fmt.Sprintf("input: %v", row.arg))
	}
}
