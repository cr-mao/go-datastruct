package singlylinkedlist

import (
	"fmt"
)

type SingleLink interface {
	GetFirstNode() *SingleLinkNode        //获取头节点
	InsertNodeFront(node *SingleLinkNode) //头部插入
	InsertNodeBack(node *SingleLinkNode)  // 尾部插入
	//在一个节点之后插入
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool
	//在一个节点之前插入
	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool
	//根据索引抓取指定位置的节点
	GetNodeAtIndex(index int) *SingleLinkNode
	// 删除一个节点
	DeleteNode(dest *SingleLinkNode) bool
	//删除指定位置的节点
	DeleteAtIndex(index int) bool
	String() string // //返回链表字符串
}

type SingLinkList struct {
	head   *SingleLinkNode
	length int // 链表的长度
}

//创建链表

func NewSingLinkList() *SingLinkList {
	head := NewSingleLinkNode(nil)
	return &SingLinkList{head: head, length: 0}
}

// 第一个元素
func (list *SingLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

//头部插入
func (list *SingLinkList) InsertNodeFront(node *SingleLinkNode) {
	// 如果没有头元素
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
	} else {
		bak := list.head
		node.pNext = bak.pNext
		list.head.pNext = node
	}
	list.length++
}

//尾部插入
func (list *SingLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
	} else {
		bak := list.head
		for bak.pNext != nil {
			bak = bak.pNext
		}
		bak.pNext = node
	}
	list.length++
}

//节点值 前追加
func (list *SingLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	var isFind = false
	bak := list.head
	for bak.pNext != nil {
		if bak.pNext.value == dest {
			isFind = true
			break
		}
		bak = bak.pNext
	}
	if isFind {
		node.pNext = bak.pNext
		bak.pNext = node
		list.length++
	}

	return isFind
}

//节点值 后追加
func (list *SingLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	var isFind = false
	bak := list.head
	for bak.pNext != nil {
		if bak.value == dest {
			isFind = true
			break
		}
		bak = bak.pNext
	}
	if isFind {
		node.pNext = bak.pNext
		bak.pNext = node
		list.length++
	}

	return isFind
}

//获得指定位置的节点 从0开始-》lenth-1  .
func (list *SingLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	}
	/*
		phead:=list.head
		for index >-1{
			phead=phead.pNext //向后循环
			index-- //向后循环过程
		}
		return phead
	*/
	var i int
	node := list.head
	for node.pNext != nil {
		node = node.pNext
		if index == i {
			//break
			return node
		}
		i++
	}
	return node
}

//删除节点
func (list *SingLinkList) DeleteNode(dest *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	node := list.head
	for node.pNext != nil {
		if node.pNext == dest {
			node.pNext = node.pNext.pNext
			list.length--
			return true
		}
		node = node.pNext
	}
	return false
}

//删除指定位置的节点
func (list *SingLinkList) DeleteAtIndex(index int) bool {
	if index > list.length-1 || index < 0 {
		return false
	}
	var i int
	node := list.head
	for node.pNext != nil {
		if index == i {
			node.pNext = node.pNext.pNext
			return true
		}
		node = node.pNext
		i++
	}
	return false
}

//打印链表
func (list *SingLinkList) String() string {
	var listString string
	node := list.head //头部节点
	for node.pNext != nil {
		listString += fmt.Sprintf("%v-->", node.pNext.value)
		node = node.pNext //循环
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}

//获得中间节点
func (list *SingLinkList) GetMid() *SingleLinkNode {
	node := list.head //头部节点
	if node == nil {
		return nil
	}
	phead1 := list.head
	phead2 := list.head
	for phead2 != nil && phead2.pNext != nil {
		phead1 = phead1.pNext
		phead2 = phead2.pNext.pNext
	}
	return phead1
}

// 链表反转
// 下一个节点 指向 前一个节点, 头节点指向 原来的最后一个节点
func (list *SingLinkList) ReverseList() {
	if list.head == nil || list.head.pNext == nil {
		return //链表为空或者链表只有一个节点
	} else {
		var pre *SingleLinkNode                   //前面节点
		var cur *SingleLinkNode = list.head.pNext //当前节点
		for cur != nil {
			curNext := cur.pNext // 后续节点
			cur.pNext = pre      //反转第一步

			pre = cur     //持续推进
			cur = curNext //持续推进
		}
		// fmt.Println(pre)
		// list.head.pNext.pNext=nil
		list.head.pNext = pre
	}
}
