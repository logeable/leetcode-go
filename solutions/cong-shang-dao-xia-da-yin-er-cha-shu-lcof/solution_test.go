package solution

import (
	"reflect"
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *types.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				root: types.MakeTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			},
			want: []int{3, 9, 20, 15, 7},
		},
		{
			name: "example 2",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
