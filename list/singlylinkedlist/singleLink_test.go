package singlylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList(t *testing.T) {
	node1 := NewSingleLinkNode(1)
	node2 := NewSingleLinkNode(2)
	node3 := NewSingleLinkNode(3)
	node4 := NewSingleLinkNode(4)
	node5 := NewSingleLinkNode(5)
	node6 := NewSingleLinkNode(6)
	node10 := NewSingleLinkNode(10)
	link := NewSingLinkList()
	t.Log(link.String())
	link.InsertNodeFront(node1)
	t.Log(link.String())
	link.InsertNodeFront(node2)
	t.Log(link.String())
	link.InsertNodeFront(node3)
	t.Log(link.String())
	t.Log("中间节点：", (link.GetMid().Value()))
	link.InsertNodeBack(node10)
	t.Log(link.String())
	t.Log(link.GetFirstNode().Value())
	link.InsertNodeValueBack(3, node4)
	t.Log(link.String())
	link.InsertNodeValueFront(2, node5)
	t.Log(link.String())
	t.Log("中间节点：", (link.GetMid().Value()))
	link.InsertNodeValueFront(3, node6)
	t.Log(link.String())

	firstNode := link.GetNodeAtIndex(0)
	assert.Equal(t, firstNode.Value(), 6, "根据索引获得节点失败")
	twoNode := link.GetNodeAtIndex(1)
	assert.Equal(t, twoNode.Value(), 3, "根据索引获得节点失败")
	//链表反转
	link.ReverseList()
	t.Log(link.String())
	link.DeleteAtIndex(0)
	t.Log(link.String())
	link.DeleteAtIndex(3)
	t.Log(link.String())
	link.DeleteNode(node10)
	t.Log(link.String())
	link.DeleteNode(node4)
	t.Log(link.String())


}
