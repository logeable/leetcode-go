package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	table = []struct {
		s        string
		expected string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
	}
)

func Test_longestPalindromeBruteForce(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, longestPalindromeBruteForce(row.s), fmt.Sprintf("input: %v", row.s))
	}
}

func Test_longestPalindrome(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, longestPalindrome(row.s), fmt.Sprintf("input: %v", row.s))
	}
}
