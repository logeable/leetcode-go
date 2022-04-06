package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				s: "abc",
			},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "example 2",
			args: args{
				s: "aab",
			},
			want: []string{"aba", "aab", "baa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check := assert.New(t)
			check.ElementsMatch(tt.want, permutation(tt.args.s))
		})
	}
}
