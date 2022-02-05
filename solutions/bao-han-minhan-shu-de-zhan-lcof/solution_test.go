package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack_1(t *testing.T) {
	check := assert.New(t)

	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	check.Equal(-3, s.Min())
	s.Pop()
	check.Equal(0, s.Top())
	check.Equal(-2, s.Min())
	s.Pop()
	check.Equal(-2, s.Top())
	check.Equal(-2, s.Min())
	s.Pop()
	check.Equal(-1, s.Top())
	check.Equal(-1, s.Min())
}
