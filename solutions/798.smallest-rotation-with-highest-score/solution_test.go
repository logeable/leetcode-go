package solution

import "testing"

func Test_bestRotation(t *testing.T) {
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
				nums: []int{2, 3, 1, 4, 0},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 3, 0, 2, 4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestRotation(tt.args.nums); got != tt.want {
				t.Errorf("bestRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
