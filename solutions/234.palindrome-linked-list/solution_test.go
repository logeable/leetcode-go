package solution

import (
	"testing"

	"github.com/logeable/leetcode-go/types"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only one",
			args: args{
				head: types.MakeListNode([]int{1}),
			},
			want: true,
		},
		{
			name: "nil",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "1221",
			args: args{
				head: types.MakeListNode([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{
			name: "12321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "123321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "1234321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 4, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "121",
			args: args{
				head: types.MakeListNode([]int{1, 2, 1}),
			},
			want: true,
		},
		{
			name: "11",
			args: args{
				head: types.MakeListNode([]int{1, 1}),
			},
			want: true,
		},
		{
			name: "12",
			args: args{
				head: types.MakeListNode([]int{1, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeUsingStack(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only one",
			args: args{
				head: types.MakeListNode([]int{1}),
			},
			want: true,
		},
		{
			name: "nil",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "1221",
			args: args{
				head: types.MakeListNode([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{
			name: "12321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "123321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "1234321",
			args: args{
				head: types.MakeListNode([]int{1, 2, 3, 4, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "121",
			args: args{
				head: types.MakeListNode([]int{1, 2, 1}),
			},
			want: true,
		},
		{
			name: "11",
			args: args{
				head: types.MakeListNode([]int{1, 1}),
			},
			want: true,
		},
		{
			name: "12",
			args: args{
				head: types.MakeListNode([]int{1, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeUsingStack(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeUsingStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
