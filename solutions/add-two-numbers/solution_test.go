package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

var (
	table = []struct {
		node1    *ListNode
		node2    *ListNode
		expected *ListNode
	}{
		{
			types.MakeListNode([]int{2, 4, 3}),
			types.MakeListNode([]int{5, 6, 4}),
			types.MakeListNode([]int{7, 0, 8}),
		},
		{
			types.MakeListNode([]int{1}),
			types.MakeListNode([]int{1}),
			types.MakeListNode([]int{2}),
		},
		{
			types.MakeListNode([]int{1}),
			types.MakeListNode([]int{9}),
			types.MakeListNode([]int{0, 1}),
		},
		{
			types.MakeListNode([]int{1, 1}),
			types.MakeListNode([]int{1, 1, 1}),
			types.MakeListNode([]int{2, 2, 1}),
		},
		{
			types.MakeListNode([]int{1, 9}),
			types.MakeListNode([]int{1, 1, 1}),
			types.MakeListNode([]int{2, 0, 2}),
		},
		{
			types.MakeListNode([]int{1, 1}),
			types.MakeListNode([]int{9}),
			types.MakeListNode([]int{0, 2}),
		},
		{
			nil,
			types.MakeListNode([]int{1}),
			types.MakeListNode([]int{1}),
		},
		{
			types.MakeListNode([]int{1}),
			nil,
			types.MakeListNode([]int{1}),
		},
		{
			nil,
			nil,
			nil,
		},
	}
)

func Test_addTwoNumbers(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, addTwoNumbers(row.node1, row.node2), fmt.Sprintf("input: %v %v", row.node1, row.node2))
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers(table[0].node1, table[0].node2)
	}
}
