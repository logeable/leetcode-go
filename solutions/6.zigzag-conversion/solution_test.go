package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dump(z [][]rune) {
	for _, row := range z {
		for _, col := range row {
			if col == 0 {
				col = ' '
			}
			fmt.Printf("%c ", col)
		}
		fmt.Println()
	}
}

func Test_zigzag(t *testing.T) {
	s := "LEETCODEISHIRING"
	numRows := 4

	fmt.Printf("row string: %q, %d rows\n", s, numRows)
	z := zigzag(s, numRows)
	fmt.Println("after zigzag:")
	dump(z)
}

func Test_convert(t *testing.T) {
	table := []struct {
		s        string
		numRows  int
		expected string
	}{
		{
			"",
			1,
			"",
		},
		{
			"abc",
			0,
			"",
		},
		{
			"A",
			1,
			"A",
		},
		{
			"ABCDEF",
			1,
			"ABCDEF",
		},
		{
			"LEETCODEISHIRING",
			4,
			"LDREOEIIECIHNTSG",
		},
		{
			"LEETCODEISHIRING",
			3,
			"LCIRETOESIIGEDHN",
		},
	}
	for _, row := range table {
		t.Logf("%q: %d", row.s, row.numRows)
		result := convert(row.s, row.numRows)
		assert.Equal(t, row.expected, result)
	}
}
