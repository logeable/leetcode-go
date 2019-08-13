package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	table = []struct {
		arg      string
		expected int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"",
			0,
		},
	}
)

func Test_lengthOfLongestSubstringBruteForce(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, lengthOfLongestSubstringBruteForce(row.arg), fmt.Sprintf("input: %v", row.arg))
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	for _, row := range table {
		assert.Equal(t, row.expected, lengthOfLongestSubstring(row.arg), fmt.Sprintf("input: %v", row.arg))
	}
}

func Benchmark_lengthOfLongestSubstringBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringBruteForce(table[0].arg)
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(table[0].arg)
	}
}
