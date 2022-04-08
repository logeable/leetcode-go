package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *types.Node) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	q := []*types.Node{root}
	cnt := 1
	for len(q) > 0 {
		var g []int
		tq := q[:cnt]
		q = q[cnt:]

		cnt = 0
		for _, n := range tq {
			g = append(g, n.Val)
			cnt += len(n.Children)
			q = append(q, n.Children...)
		}
		ans = append(ans, g)
	}
	return ans
}
