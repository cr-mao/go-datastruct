package arraylist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList()
	list.Append(2)
	list.Append(3)
	list.Append("222")
	t.Log(list)
	t.Log(list.TheSize)
	assert.Equal(t, list.TheSize, 3, "the size should be equal")
	index0, _ := list.Get(0)
	assert.Equal(t, 2, index0, "the size should be equal")
	var p List = NewArrayList()
	for i := 0; i < 20; i++ {
		p.Append(i)
		t.Log(p)
	}
	p.Clear()
	t.Log(p)
}
