package solution

import (
	"testing"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		ok   bool
	}{
		{
			name: "example 1",
			args: args{
				rolls: []int{3, 2, 4, 3},
				mean:  4,
				n:     2,
			},
			ok: true,
		},
		{
			name: "example 2",
			args: args{
				rolls: []int{1, 5, 6},
				mean:  3,
				n:     4,
			},
			ok: true,
		},
		{
			name: "example 3",
			args: args{
				rolls: []int{1, 2, 3, 4},
				mean:  6,
				n:     4,
			},
			ok: false,
		},
		{
			name: "example 4",
			args: args{
				rolls: []int{1},
				mean:  3,
				n:     1,
			},
			ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rolls := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n)
			if !tt.ok && rolls != nil {
				t.Error("expect no rolls, but got", rolls)
			}
			if tt.ok && !check(tt.args.rolls, rolls, tt.args.mean) {
				t.Error("rolls is invalid", rolls)
			}
		})
	}
}

func check(m, n []int, mean int) bool {
	total := 0
	for _, x := range m {
		total += x
	}
	for _, x := range n {
		total += x
	}
	return total == mean*(len(m)+len(n))
}
