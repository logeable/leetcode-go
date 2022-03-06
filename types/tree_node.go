package types

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	var buf bytes.Buffer
	dump(&buf, t, 0)
	return buf.String()
}

func dump(w io.Writer, t *TreeNode, space int) {
	if t == nil {
		return
	}
	const (
		HPad = 4
		VPad = 2
	)

	dump(w, t.Right, space+HPad)

	w.Write([]byte(strings.Repeat(" ", space)))
	w.Write([]byte(strconv.Itoa(t.Val)))
	w.Write([]byte(strings.Repeat("\n", VPad)))
	dump(w, t.Left, space+HPad)
}

func MakeTreeNode(data []interface{}) *TreeNode {
	originData := data
	if len(data) == 0 || data[0] == nil {
		return nil
	}

	root := &TreeNode{data[0].(int), nil, nil}
	data = data[1:]

	var queue []*TreeNode
	queue = append(queue, root)
	// flag: true for left, false for right
	flag := true
	for len(queue) != 0 {
		if len(data) == 0 {
			break
		}
		var d interface{}
		d, data = data[0], data[1:]
		q := queue[0]
		if d == nil {
			if !flag {
				queue = queue[1:]
			}
			flag = !flag
			continue
		}
		node := &TreeNode{d.(int), nil, nil}
		queue = append(queue, node)
		if flag {
			q.Left = node
		} else {
			q.Right = node
			queue = queue[1:]
		}
		flag = !flag
	}
	if len(data) > 0 {
		panic(fmt.Errorf("invalid data: %v", originData))
	}

	return root
}

func TreeEqual(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && TreeEqual(t1.Left, t2.Left) && TreeEqual(t1.Right, t2.Right)
	} else {
		return false
	}
}
