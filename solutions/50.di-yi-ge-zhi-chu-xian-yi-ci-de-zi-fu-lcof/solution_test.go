package solution

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "example 1",
			args: args{
				s: "abaccdeff",
			},
			want: 'b',
		},
		{
			name: "example 2",
			args: args{
				s: "",
			},
			want: ' ',
		},
		{
			name: "example 3",
			args: args{
				s: "cc",
			},
			want: ' ',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
