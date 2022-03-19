package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				root: types.MakeTreeNode([]interface{}{1, 2, 3, 4}),
			},
			want: "1(2(4))(3)",
		},
		{
			name: "example 2",
			args: args{
				root: types.MakeTreeNode([]interface{}{1, 2, 3, nil, 4}),
			},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
