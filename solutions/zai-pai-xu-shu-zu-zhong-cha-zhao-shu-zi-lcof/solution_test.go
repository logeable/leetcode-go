package solution

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: 0,
		},
		{
			name: "one element",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "one element",
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: 0,
		},
		{
			name: "empty",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
