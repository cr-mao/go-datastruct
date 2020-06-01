package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack(t *testing.T) {
	arraystack := NewStack()
	arraystack.Push(1)
	arraystack.Push(2)
	v := arraystack.Pop()
	t.Log(v)
	assert.Equal(t, 2, v)
	v = arraystack.Pop()
	t.Log(v)
	assert.Equal(t, 1, v)
}
