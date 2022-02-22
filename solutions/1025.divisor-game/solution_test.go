package solution

import "testing"

func Test_divisorGame(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				n: 2,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				n: 3,
			},
			want: false,
		},
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.n); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
