package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				root: types.MakeTreeNode([]interface{}{1, 2, 2, 3, 4, 4, 3}),
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				root: types.MakeTreeNode([]interface{}{1, 2, 2, nil, 3, nil, 3}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
