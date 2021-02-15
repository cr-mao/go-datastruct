package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkStack(t *testing.T) {
	arraystack := NewLinkStack()
	arraystack.Push(1)
	arraystack.Push(2)
	v := arraystack.Pop()
	t.Log(v)
	assert.Equal(t, 2, v)
	v = arraystack.Pop()
	t.Log(v)
	assert.Equal(t, 1, v)

	for i := 0; i < 10000000; i++ {
		arraystack.Push(i)
	}

	for data := arraystack.Pop(); data != nil; data = arraystack.Pop() {
		fmt.Println(data)
	}
}
