package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: []int{1, 3, 4, 2, 2}},
			want: 2,
		},
		{
			name: "2",
			args: args{nums: []int{3, 1, 3, 4, 2}},
			want: 3,
		},
		{
			name: "3",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "4",
			args: args{nums: []int{1, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := assert.New(t)
			nums := make([]int, len(tt.args.nums))
			copy(nums, tt.args.nums)
			got := findDuplicate(nums)
			check.Equal(tt.want, got)
			check.Equal(tt.args.nums, nums)
		})
	}
}
