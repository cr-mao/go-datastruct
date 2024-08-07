/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/30
**/
package tree

import (
	list2 "container/list"
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	/*
			    4(node1)
			2(node2)                   6(node3)
		1(node4)	  3(node5)      5(node6)  7 (node7)
			                      17(node8)
	*/
	bst := NewBinaryTree()
	node1 := &Node{4, nil, nil}
	node2 := &Node{2, nil, nil}
	node3 := &Node{6, nil, nil}
	node4 := &Node{1, nil, nil}
	node5 := &Node{3, nil, nil}
	node6 := &Node{5, nil, nil}
	node7 := &Node{7, nil, nil}
	//node8 := &Node{17, nil, nil}
	bst.Root = node1
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	//node6.Left = node8
	bst.Size = 7

	fmt.Println(bst.FindMax(), "max")
	fmt.Println(bst.FindMin(), "min")
	fmt.Println("中序----------------------")
	bst.InOrder()
	fmt.Println("前序----------------------")
	bst.PreOrder()
	fmt.Println("后序----------------------")
	bst.PostOrder()
	fmt.Println("last ----------------------")
	fmt.Println(bst)
	//// 移除最大节点
	//bst.RemoveMax()
	//fmt.Println(bst)
	//// 移除最小节点
	//bst.RemoveMin()
	//fmt.Println(bst)
	// 广度遍历
	fmt.Println("广度遍历----------------------")
	bst.Levelshow()
	// 深度遍历
	fmt.Println("深度遍历----------------------")
	bst.Stackshow(bst.Root)
	// 最小公共祖先
	fmt.Println(bst.FindlowerstAncestor(bst.Root, node2, node6).Data)

}

func TestList(t *testing.T) {
	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)
	value := list.Back().Value.(int)
	fmt.Println(value)
}
