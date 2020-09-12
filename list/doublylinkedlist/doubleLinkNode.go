package doublylinkedlist

type DoubleLinkNode struct {
	value interface{}
	prev  *DoubleLinkNode //上一个节点
	next  *DoubleLinkNode //下一个节点
}

func (node *DoubleLinkNode) Value() interface{} {
	return node.value
}

func NewDoubleLinkNode(value interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{value, nil, nil}
}

func (node *DoubleLinkNode) Prev() *DoubleLinkNode {
	return node.prev
}

func (node *DoubleLinkNode) Next() *DoubleLinkNode {
	return node.next
}
