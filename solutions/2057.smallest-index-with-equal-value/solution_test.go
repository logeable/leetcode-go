package solution

import "testing"

func Test_smallestEqual(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEqual(tt.args.nums); got != tt.want {
				t.Errorf("smallestEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
