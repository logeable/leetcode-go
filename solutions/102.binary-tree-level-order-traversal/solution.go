package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *types.TreeNode) [][]int {
	heightMap := make(map[*types.TreeNode]int)
	queue := []*types.TreeNode{root}
	heightMap[root] = 0

	lastLevel := -1
	var result [][]int
	for len(queue) != 0 {
		var d *types.TreeNode
		d, queue = queue[0], queue[1:]
		if d == nil {
			continue
		}
		h := heightMap[d]
		if heightMap[d] > lastLevel {
			lastLevel = h
			result = append(result, make([]int, 0))
		}
		result[h] = append(result[h], d.Val)

		if d.Left != nil {
			queue = append(queue, d.Left)
			heightMap[d.Left] = h + 1
		}
		if d.Right != nil {
			queue = append(queue, d.Right)
			heightMap[d.Right] = h + 1
		}
	}

	return result
}
