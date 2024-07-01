/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package tree

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node // 根节点
	Size int   // 树中节点的个数
}

// 新建一个二叉树
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
		Size: 0,
	}
}

// 获取二叉树大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

// 判断是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

// 根节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

// 插入节点
func (bst *BinaryTree) add(n *Node, data int) *Node {
	if n == nil {
		bst.Size++
		return &Node{Data: data, Left: nil, Right: nil}
	}
	if data < n.Data {
		n.Left = bst.add(n.Left, data) //比我小，左边
	} else if data > n.Data {
		n.Right = bst.add(n.Right, data)
	}
	return n
}

// 二分查找 查找节点
func (bst *BinaryTree) isIn(n *Node, data int) bool {
	if n == nil {
		return false
	}
	if data == n.Data {
		return true
	}
	if data < n.Data {
		return bst.isIn(n.Left, data)
	}
	return bst.isIn(n.Right, data)
}

// 最大元素
func (bsg *BinaryTree) FindMax() int {
	if bsg.Root == nil {
		panic("根节点不存在")
	}
	return bsg.findMax(bsg.Root).Data
}

// 最大元素
func (bst *BinaryTree) findMax(n *Node) *Node {
	if n.Right == nil {
		return n
	}
	return bst.findMax(n.Right)
}

// 最小元素
func (bsg *BinaryTree) FindMin() int {
	if bsg.Root == nil {
		panic("根节点不存在")
	}
	return bsg.findMin(bsg.Root).Data
}

// 最小元素
func (bst *BinaryTree) findMin(n *Node) *Node {
	if n.Left == nil {
		return n
	}
	return bst.findMin(n.Left)
}

// 前序遍历
func (bst *BinaryTree) PreOrder() {
	bst.preOrder(bst.Root)
}
func (bst *BinaryTree) preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Data)
	bst.preOrder(n.Left)
	bst.preOrder(n.Right)
}

// 中序遍历
func (bst *BinaryTree) InOrder() {
	bst.inOrder(bst.Root)
}
func (bst *BinaryTree) inOrder(n *Node) {
	if n == nil {
		return
	}
	bst.preOrder(n.Left)
	fmt.Println(n.Data)
	bst.preOrder(n.Right)
}

// 后序遍历
func (bst *BinaryTree) PostOrder() {
	bst.postOrder(bst.Root)
}
func (bst *BinaryTree) postOrder(n *Node) {
	if n == nil {
		return
	}
	bst.preOrder(n.Left)
	bst.preOrder(n.Right)
	fmt.Println(n.Data)
}

// 打印数形 中序遍历
func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer                     //保存字符串
	bst.GenerateBSTstring(bst.Root, 0, &buffer) //调用函数实现遍历
	return buffer.String()
}

func (bst *BinaryTree) GenerateBSTstring(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		//buffer.WriteString(bst.GenerateDepthstring(depth)+"nil\n")//空节点
		return
	}
	//写入字符串，保存树的深度
	bst.GenerateBSTstring(node.Left, depth+1, buffer)
	buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Data) + "\n")
	bst.GenerateBSTstring(node.Right, depth+1, buffer)
}

func (bst *BinaryTree) GenerateDepthstring(depth int) string {
	var buffer bytes.Buffer //保存字符串
	for i := 0; i < depth; i++ {
		buffer.WriteString("--") //深度为0，1-- 2----
	}
	return buffer.String()
}

// 移除最大节点
func (bst *BinaryTree) RemoveMax() int {
	max := bst.FindMax()
	bst.Root = bst.removeMax(bst.Root)
	return max
}
func (bst *BinaryTree) removeMax(n *Node) *Node {
	if n.Right == nil {
		leftNode := n.Left
		bst.Size--
		return leftNode
	}
	n.Right = bst.removeMax(n.Right)
	return n
}

// 移除最小节点
func (bst *BinaryTree) RemoveMin() int {
	min := bst.FindMin()
	bst.Root = bst.removeMin(bst.Root)
	return min
}
func (bst *BinaryTree) removeMin(n *Node) *Node {
	if n.Left == nil {
		rightNode := n.Right
		bst.Size--
		return rightNode
	}
	n.Left = bst.removeMin(n.Left)
	return n
}
