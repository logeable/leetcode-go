package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "functional test",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 4, 5}),
			},
			want: types.MakeListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "zero node",
			args: args{
				head: types.MakeListNode([]int{}),
			},
			want: types.MakeListNode([]int{}),
		},
		{
			name: "one node",
			args: args{
				head: types.MakeListNode([]int{1}),
			},
			want: types.MakeListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !types.ListEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "functional test",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 4, 5}),
			},
			want: types.MakeListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "zero node",
			args: args{
				head: types.MakeListNode([]int{}),
			},
			want: types.MakeListNode([]int{}),
		},
		{
			name: "one node",
			args: args{
				head: types.MakeListNode([]int{1}),
			},
			want: types.MakeListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListRecursive(tt.args.head); !types.ListEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
