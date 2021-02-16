package queue

import "testing"

func TestLinkQueue(t *testing.T) {
	myq := NewLinkQueue()
	for i := 1; i < 10000000; i++ {
		myq.Enqueue(i)
	}
	for data := myq.Dequeue(); data != nil; data = myq.Dequeue() {
		t.Log(data)
	}
}
