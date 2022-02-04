package solution

import (
	"reflect"
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				head: types.MakeListNode([]int{1, 3, 2}),
			},
			want: []int{2, 3, 1},
		},
		{
			name: "2",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
