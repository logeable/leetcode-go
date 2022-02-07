package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *types.ListNode) *types.ListNode {
	p := head
	prev := (*types.ListNode)(nil)
	for p != nil {
		n := p.Next
		p.Next = prev
		prev = p
		p = n
	}

	return prev
}
