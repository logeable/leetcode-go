package solution

import "testing"

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				colors: "AAABABB",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				colors: "AAAABBBB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
