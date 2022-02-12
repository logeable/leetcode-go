package solution

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{n: 3},
			want: 0,
		},
		{
			name: "example 2",
			args: args{n: 5},
			want: 1,
		},
		{
			name: "example 3",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "40",
			args: args{n: 40},
			want: 9,
		},
		{
			name: "10000",
			args: args{n: 10000},
			want: 2499,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
