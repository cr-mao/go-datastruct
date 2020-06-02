package arraylist

import (
	"testing"
)

func TestArrayListIteratorStack(t *testing.T) {
	mystack := NewArrayListStackIterator()
	mystack.Push(1)
	t.Log("入栈", 1)
	mystack.Push(2)
	t.Log("入栈", 2)
	mystack.Push(3)
	t.Log("入栈", 3)
	mystack.Push(4)
	t.Log("入栈", 4)
	t.Log(mystack.myarray)
	t.Log("出栈", mystack.Pop())
	t.Log("出栈", mystack.Pop())
	t.Log("出栈", mystack.Pop())
	t.Log("出栈", mystack.Pop())
	mystack.Push(1)
	t.Log("入栈", 1)
	mystack.Push(2)
	t.Log("入栈", 2)
	mystack.Push(3)
	t.Log("入栈", 3)
	mystack.Push(4)
	t.Log("入栈", 4)
	t.Log("遍历栈数组")
	for it := mystack.Myit; it.HasNext(); {
		val, _ := it.Next()
		t.Log("遍历出", val)
	}
}
