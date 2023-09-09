package solution

import (
	"reflect"
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "112",
			args: args{head: types.MakeListNode([]int{1, 1, 2})},
			want: types.MakeListNode([]int{1, 2}),
		},
		{
			name: "11233",
			args: args{head: types.MakeListNode([]int{1, 1, 2, 3, 3})},
			want: types.MakeListNode([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
