package arraylist

import (
	"testing"
)

func TestArrayListStack(t *testing.T) {
	mystack := NewArrayListStack()
	mystack.Push(1)
	t.Log("入栈",1)
	mystack.Push(2)
	t.Log("入栈",2)
	mystack.Push(3)
	t.Log("入栈",3)
	mystack.Push(4)
	t.Log("入栈",4)
	t.Log("出栈",mystack.Pop())
	t.Log("出栈",mystack.Pop())
	t.Log("出栈",mystack.Pop())
	t.Log("出栈",mystack.Pop())
	mystack.Push(1)
	t.Log("入栈",1)
	t.Log("出栈",mystack.Pop())
	mystack.Push(2)
	t.Log("入栈",2)
	t.Log("出栈",mystack.Pop())
	mystack.Push(3)
	t.Log("入栈",3)
	t.Log("出栈",mystack.Pop())
	mystack.Push(4)
	t.Log("入栈",4)
	t.Log("出栈",mystack.Pop())

}
