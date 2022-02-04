package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *types.ListNode) []int {
	if head == nil {
		return nil
	}

	r := reversePrint(head.Next)
	r = append(r, head.Val)

	return r
}
