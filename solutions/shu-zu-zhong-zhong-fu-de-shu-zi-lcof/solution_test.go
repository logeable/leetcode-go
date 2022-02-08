package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := assert.New(t)
			got := findRepeatNumber(tt.args.nums)
			check.Contains(tt.want, got)
		})
	}
}
