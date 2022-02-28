package solution

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 3,
			},
			want: []int{3},
		},
		{
			name: "example 3",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 5,
			},
			want: []int{4},
		},
		{
			name: "example 4",
			args: args{
				nums:   []int{1, 2, 5, 2, 3},
				target: 4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("targetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
