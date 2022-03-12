package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for a Node.
  * type Node struct {
	   *     Val int
	    *     Children []*Node
		 * }
*/

func postorder(root *types.Node) []int {
	if root == nil {
		return nil
	}

	var ret []int
	for _, child := range root.Children {
		ret = append(ret, postorder(child)...)
	}
	ret = append(ret, root.Val)
	return ret
}
