package solution

import (
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "example 2",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := combinationSum(tt.args.candidates, tt.args.target)
			if len(tt.want) != len(ans) {
				t.Errorf("length unexpect")
			}
			for _, v := range ans {
				total := 0
				for _, t := range v {
					total += t
				}
				if total != tt.args.target {
					t.Errorf("sum invalid")
				}
			}
		})
	}
}
