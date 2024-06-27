/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package tree

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
