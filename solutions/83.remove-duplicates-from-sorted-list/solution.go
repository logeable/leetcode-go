package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = types.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	cur := head.Next
	for cur != nil {
		if cur.Val != p.Val {
			p.Next = cur
			p = cur
		}
		cur = cur.Next
	}
	p.Next = nil
	return head
}
