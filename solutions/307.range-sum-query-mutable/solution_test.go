package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray(t *testing.T) {
	check := assert.New(t)

	arr := Constructor([]int{1, 3, 5})
	check.Equal(9, arr.SumRange(0, 2))
	arr.Update(1, 2)
	check.Equal(8, arr.SumRange(0, 2))
}
