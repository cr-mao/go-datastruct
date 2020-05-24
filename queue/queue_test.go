package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	myq := NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	if myq.Size() != 3 {
		t.Error("enqueue error")
	}
	front := myq.Front()
	if front != 1 {
		t.Error("front error")
	}
	end := myq.End()
	if end != 3 {
		t.Error("end error")
	}
	data := myq.DeQueue()
	if data != 1 {
		t.Error("dequeue error")
	}
	myq.DeQueue()
	myq.Clear()
	if myq.Size() != 0 {
		t.Error("clear error")
	}
	if !myq.IsEmpty() {
		t.Error("IsEmpty error")
	}
}
