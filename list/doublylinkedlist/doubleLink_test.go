package doublylinkedlist

import (
	"testing"
)


func TestDoubleLink(t *testing.T) {
	dobleLink :=NewDoubleLinkList()

	node1 :=NewDoubleLinkNode(1)
	node2 :=NewDoubleLinkNode(2)
	node3 :=NewDoubleLinkNode(3)
	node4 :=NewDoubleLinkNode(4)
	node5 :=NewDoubleLinkNode(5)
	node6 :=NewDoubleLinkNode(6)
	node7 :=NewDoubleLinkNode(7)
	node8 :=NewDoubleLinkNode(8)
	node9 :=NewDoubleLinkNode(9)
	node10 :=NewDoubleLinkNode(10)
	dobleLink.InsertHead(node1)
	dobleLink.InsertHead(node2)
	dobleLink.InsertHead(node3)
	dobleLink.InsertHead(node4)
	t.Log(dobleLink.String())
	dobleLink.InsertBack(node5)
	dobleLink.InsertBack(node6)	
	t.Log(dobleLink.String())
	dobleLink.DeleteNode(node5)
	t.Log(dobleLink.String())
	dobleLink.DeleteNode(node6)
	t.Log(dobleLink.String())
	t.Log("获得第一个节点值",dobleLink.GetFirstNode().Value())
	dobleLink.InsertNodeBack(node4,node7)
	t.Log(dobleLink.String())
	dobleLink.InsertNodeHead(node4,node8)
	t.Log(dobleLink.String())
	dobleLink.InsertValueBack(8,node9)
	t.Log(dobleLink.String())
	dobleLink.InsertValueHead(8,node10)
	t.Log(dobleLink.String())
	dobleLink.DeleteNodeAtindex(0)
	t.Log(dobleLink.String())
}