package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *types.ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: types.MakeListNode([]int{4, 5, 1, 9}),
				val:  5,
			},
			want: types.MakeListNode([]int{4, 1, 9}),
		},
		{
			name: "example 2",
			args: args{
				head: types.MakeListNode([]int{4, 5, 1, 9}),
				val:  1,
			},
			want: types.MakeListNode([]int{4, 5, 9}),
		},
		{
			name: "example 3",
			args: args{
				head: types.MakeListNode([]int{4, 5, 1, 9}),
				val:  4,
			},
			want: types.MakeListNode([]int{5, 1, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !types.ListEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
