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
