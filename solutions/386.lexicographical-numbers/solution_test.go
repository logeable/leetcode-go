package solution

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				n: 13,
			},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "example 2",
			args: args{
				n: 2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
