package arraylist

import (
	"testing"
)

func TestArrayListIterator(t *testing.T) {
	list := NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("222")
	it := list.Iterator()
	for it.HasNext() {
		val, _ := it.Next()
		t.Log(val)
	}
}
