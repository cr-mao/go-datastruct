package doublylinkedlist

import (
	"fmt"
	"strings"
)

type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

//新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	node := NewDoubleLinkNode(nil)
	return &DoubleLinkList{node, 0}
}

//获得链表长度
func (dlist *DoubleLinkList) GetLength() int {
	return dlist.length
}

//返回第一个节点 （数据值不为nil的节点）
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

// 头部插入 节点

func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := dlist.head //头节点
	// 没有第一个节点
	if phead.next == nil {
		node.next = nil
		phead.next = node
		node.prev = phead
	} else {
		//已经有第一个节点
		phead.next.prev = node //标记上一个节点
		node.next = phead.next //下一个节点
		phead.next = node      //标记头部节点的下一个节点
		node.prev = phead      //node上一个节点 指向头节点
	}
	dlist.length++
}

//尾部插入节点
func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head
	for phead.next != nil {
		phead = phead.next
	}
	//phead 为最后一个节点
	phead.next = node
	// node.next=nil 可以省略 node 默认next 就是nil
	node.prev = phead
	dlist.length++
}

// 打印链表
func (dlist *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil {
		listString1 += fmt.Sprintf("%v-->", phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	// 此时 phead 已经是最后一个节点了
	for phead != dlist.head {
		//反向循环
		listString2 += fmt.Sprintf("<--%v", phead.value)
		phead = phead.prev
	}
	listString1 += fmt.Sprintf("nil")
	return "\n" + listString1 + listString2 + "\n" //打印链表字符串
}

//指定节点后面插入
func (dlist *DoubleLinkList) InsertNodeBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest { //循环查找
		phead = phead.next
	}
	//  value  cur     prev   next
	//   1    400100  40000  40200
	//   2    40200   400100  40300
	//   3     40300  40200   nil
	// node   400150
	//目标节点是 2 后面插入
	if phead.next == dest {
		//看已经是不是最后了 是最后了 只有3步 ，否则是4步
		if phead.next.next != nil {
			phead.next.next.prev = node //400150
		}
		node.next = phead.next.next //400300
		phead.next.next = node      // 400150
		node.prev = phead.next      //  400200
		dlist.length++
		return true
	} else {
		return false
	}
}

//指定节点前面 插入
func (dlist *DoubleLinkList) InsertNodeHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {

	phead := dlist.head
	for phead.next != nil && phead.next != dest {
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next != nil { //这个应该去掉
			phead.next.prev = node
		}
		node.next = phead.next
		node.prev = phead
		phead.next = node
		dlist.length++
		return true
	} else {
		return false
	}
}

//根据值 后插入
func (dlist *DoubleLinkList) InsertValueBack(dest interface{}, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next.value != dest {
		phead = phead.next
	}
	if phead.next.value == dest {
		if phead.next.next != nil {
			phead.next.next.prev = node
		}
		node.next = phead.next.next
		phead.next.next = node
		node.prev = phead.next
		dlist.length++
		return true
	} else {
		return false
	}
}

//根据值 前插入
func(dlist* DoubleLinkList) InsertValueHead(dest  interface{},node *DoubleLinkNode )bool{
	phead:=dlist.head
	for phead.next!=nil && phead.next.value!=dest{  //循环查找
		phead=phead.next
	}
	if phead.next.value==dest{
		if phead.next!=nil{
			phead.next.prev=node
		}
		node.next=phead.next
		node.prev=phead    
		phead.next=node  
		dlist.length++
		return true
	}else{
		return false
	}
}

// 根据索引获得节点  0 代表 第一个节点，头节点的值是nil
func (dlist *DoubleLinkList) GetNodeAtindex(index int) *DoubleLinkNode {
	if index > dlist.length-1 || index < 0 {
		return nil
	}
	phead := dlist.head
	for index > -1 {
		phead = phead.next
		index-- //计算位置
	}
	return phead

	// var i int
	// node := dlist.head
	// for node.next != nil {
	// 	node = node.next
	// 	if index == i {
	// 		//break
	// 		return node
	// 	}
	// 	i++
	// }
	// return node
}

//删除指定节点
func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	if node == nil {
		return false
	} else {
		phead := dlist.head
		for phead.next != nil && phead.next != node {
			phead = phead.next //循环查找
		}
		if phead.next == node {
			if phead.next.next != nil {
				phead.next.next.prev = phead //设置pre
			}
			phead.next = phead.next.next //设置next
			dlist.length--
			return true
		} else {
			return false
		}
	}
}

//删除指定位置节点
func (dlist *DoubleLinkList) DeleteNodeAtindex(index int) bool {
	if index > dlist.length-1 || index < 0 {
		return false
	}
	phead := dlist.head
	for index > 0 {
		phead = phead.next
		index-- //计算位置
	}

	if phead.next.next != nil {
		phead.next.next.prev = phead //设置pre
	}
	phead.next = phead.next.next //设置next
	dlist.length--
	return true
}

func (dlist *DoubleLinkList) FindString(inputstr string) {
	phead := dlist.head.next
	for phead.next != nil {
		//正向循环
		if strings.Contains(phead.value.(string), inputstr) {
			fmt.Println(phead.value.(string))
		}
		phead = phead.next
	}
}
