package singlylinkedlist

type SingleLinkNode struct {
	value interface{}
	pNext *SingleLinkNode
}

//构造一个节点
func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	var node = new(SingleLinkNode)
	node.value = data
	return node
}

// 返回值
func (node *SingleLinkNode) Value() interface{} {
	return node.value
}

// 返回下一个节点
func (node *SingleLinkNode) Pnext() *SingleLinkNode {
	return node.pNext
}
