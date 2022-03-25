package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *types.ListNode, val int) *types.ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	p, q := head, head.Next
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
			break
		}
		p = q
		q = q.Next
	}

	return head
}
