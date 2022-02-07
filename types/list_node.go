package types

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "()"
	}
	var buf bytes.Buffer

	var data []string
	head := n
	for head != nil {
		data = append(data, strconv.Itoa(head.Val))
		head = head.Next
	}

	buf.WriteString(fmt.Sprintf("(%s)", strings.Join(data, " -> ")))
	return buf.String()
}

func MakeListNode(data []int) *ListNode {
	var head, p *ListNode

	for _, x := range data {
		cur := &ListNode{
			x,
			nil,
		}
		if head == nil {
			head = cur
		} else {
			p.Next = cur
		}
		p = cur
	}
	return head
}

func ListEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == l2 && l1 == nil
}
