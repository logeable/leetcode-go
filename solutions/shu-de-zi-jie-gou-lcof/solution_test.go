package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *types.TreeNode
		B *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				A: types.MakeTreeNode([]interface{}{1, 2, 3}),
				B: types.MakeTreeNode([]interface{}{3, 1}),
			},
			want: false,
		},
		{
			name: "example 2",
			args: args{
				A: types.MakeTreeNode([]interface{}{3, 4, 5, 1}),
				B: types.MakeTreeNode([]interface{}{4, 1}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
