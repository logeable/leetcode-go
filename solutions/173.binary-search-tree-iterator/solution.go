package solution

import (
	"github.com/logeable/leetcode-go/types"
)

type TreeNode = types.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{nil, root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stack = append(this.stack, this.cur)
		this.cur = this.cur.Left
	}
	var top *TreeNode
	top, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	if top.Right != nil {
		this.cur = top.Right
	}
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
