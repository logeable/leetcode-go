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

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	f := head
	s := head
	for s.Next != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	if s.Next == nil {
		return true
	}
	s = s.Next
	var prev *ListNode
	for s != nil {
		next := s.Next
		s.Next = prev
		prev = s
		s = next
	}

	p := head
	q := prev
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func isPalindromeUsingStack(head *ListNode) bool {
	if head == nil {
		return false
	}
	f := head
	s := head
	for s.Next != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	var stack []*ListNode
	for s.Next != nil {
		stack = append(stack, s.Next)
		s = s.Next
	}
	p := head
	var last *ListNode
	for len(stack) != 0 {
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if last.Val != p.Val {
			return false
		}
		p = p.Next
	}

	return true
}
