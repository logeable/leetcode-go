package solution

import (
	"github.com/logeable/leetcode-go/types"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = types.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var c int

	var head, p *ListNode
	for l1 != nil || l2 != nil {
		t := c
		if l1 != nil {
			t += l1.Val
		}
		if l2 != nil {
			t += l2.Val
		}

		c = t / 10
		v := t % 10

		tmp := &ListNode{
			v,
			nil,
		}

		if head == nil {
			head = tmp
		} else {
			p.Next = tmp
		}

		p = tmp

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if c != 0 {
		p.Next = &ListNode{
			c,
			nil,
		}
	}

	return head
}
