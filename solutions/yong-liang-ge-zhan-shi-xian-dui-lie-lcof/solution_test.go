package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQueue_1(t *testing.T) {
	check := assert.New(t)

	obj := Constructor()
	obj.AppendTail(3)
	check.Equal(3, obj.DeleteHead())
	check.Equal(-1, obj.DeleteHead())
}

func TestCQueue_2(t *testing.T) {
	check := assert.New(t)

	obj := Constructor()
	check.Equal(-1, obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	check.Equal(5, obj.DeleteHead())
	check.Equal(2, obj.DeleteHead())
}
