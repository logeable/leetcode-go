package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_mirrorTree(t *testing.T) {
	type args struct {
		root *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *types.TreeNode
	}{
		{
			name: "example 1",
			args: args{
				root: types.MakeTreeNode([]interface{}{4, 2, 7, 1, 3, 6, 9}),
			},
			want: types.MakeTreeNode([]interface{}{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree(tt.args.root); !types.TreeEqual(tt.want, got) {
				t.Errorf("mirrorTree() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
