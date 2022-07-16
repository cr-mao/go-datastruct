package main

type Node struct {
	Data  int   // 数据
	Left  *Node //左边节点
	Right *Node //右边节点
}

// 二叉树
type BinaryTree struct {
	Root *Node //根节点
	Size int   //数据的数量
}

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

//是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

//从根节点开始插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

//插入节点
func (bst *BinaryTree) add(n *Node, data int) *Node {
	//到了尽头
	if n == nil {
		bst.Size++
		return &Node{
			Data:  data,
			Left:  nil,
			Right: nil,
		}
	} else {
		if data < n.Data {
			n.Left = bst.add(n.Left, data)
		} else if data > n.Data {
			n.Right = bst.add(n.Right, data)
		}
		return n
	}
}

//判断数据是否存在
func (bst *BinaryTree) Isin(data int) bool {
	return bst.isin(bst.Root, data)
}

func (bst *BinaryTree) isin(n *Node, data int) bool {
	if n == nil {
		return false //树是空树，招不到
	}
	if data == n.Data {
		return true
	} else if data < n.Data {
		return bst.isin(n.Left, data)
	} else {
		return bst.isin(n.Right, data)
	}
}

//取二叉树的最大值
func (bst *BinaryTree) FindMax() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmax(bst.Root).Data
}

func (bst *BinaryTree) findmax(n *Node) *Node {
	if n.Right != nil {
		return bst.findmax(n.Right)
	}
	return n
}

//取二叉树的 最小值
func (bst *BinaryTree) FindMin() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.findmin(bst.Root).Data
}

func (bst *BinaryTree) findmin(n *Node) *Node {
	if n.Left != nil {
		return bst.findmin(n.Left)
	}
	return n
}
func main() {

}
