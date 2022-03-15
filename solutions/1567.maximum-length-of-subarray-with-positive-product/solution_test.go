package solution

import "testing"

func Test_getMaxLen(t *testing.T) {
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
				nums: []int{1, -2, -3, 4},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, -2, -3, -4},
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{-1, -2, -3, 0, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxLen(tt.args.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
