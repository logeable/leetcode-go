package solution

import (
	"testing"
)

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 13,
				k: 2,
			},
			want: 10,
		},
		{
			name: "example 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				n: 10,
				k: 3,
			},
			want: 2,
		},
		{
			name: "example 4",
			args: args{
				n: 100,
				k: 90,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
