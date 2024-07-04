/**
* @Author: maozhongyu
* @Desc: 二叉树
* @Date: 2024/6/27
**/
package tree

import (
	"bytes"
	"container/list"
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

// 移除节点
func (bst *BinaryTree) Remove(data int) {
	bst.Root = bst.remove(bst.Root, data) //删除数据
}

func (bst *BinaryTree) remove(n *Node, data int) *Node {
	if n == nil {
		return nil //节点为空，不需要干活
	}
	if data < n.Data {
		n.Left = bst.remove(n.Left, data) //递归左边
		return n
	} else if data > n.Data {
		n.Right = bst.remove(n.Right, data) //递归右边
		return n
	} else {
		//处理左边为空
		if n.Left == nil {
			rightNode := n.Right //备份右边节点
			n.Right = nil
			bst.Size-- //删除
			return rightNode
		}
		//处理右边为空
		if n.Right == nil {
			leftNode := n.Left //备份右边节点
			n.Left = nil       //处理节点返回
			bst.Size--         //删除
			return leftNode
		}
		//左右节点都不为空
		oknode := bst.findMin(n.Right)        //找出比我小的节点顶替我， 右边的最小 也是比自己大的。
		oknode.Right = bst.removeMin(n.Right) // 移除  右边 自身这个小的
		oknode.Left = n.Left                  // 小节点左边，连接原来的左边。
		n.Left = nil                          //删除的清空
		n.Right = nil
		return oknode //实现删除

	}
}

// 非递归中序
func (bst *BinaryTree) InOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst) //首先左边压入栈中
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()           //挨个取出节点
			mybst = v.Value.(*Node)       //实例化
			res = append(res, mybst.Data) //压入数据
			mybst = mybst.Right           //追加
			mystack.Remove(v)             //删除
		}
	}
	return res
}

// 非递归前序
func (bst *BinaryTree) PreOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			res = append(res, mybst.Data) //压入数据
			mystack.PushBack(mybst)       //首先左边压入栈中
			mybst = mybst.Left
		}
		if mystack.Len() != 0 {
			v := mystack.Back()     //挨个取出节点
			mybst = v.Value.(*Node) //实例化
			//res=append(res,mybst.Data)//压入数据
			mybst = mybst.Right //追加
			mystack.Remove(v)   //删除
		}
	}
	return res
}

// 后序非递归
func (bst *BinaryTree) PostOrderNoRecursion() []int {
	mybst := bst.Root     //备份二叉树
	mystack := list.New() //生成一个栈
	res := make([]int, 0) //生成数组，容纳中序的数据
	var PreVisited *Node  //提前访问的节点 , 后续 要记录下 之前的节点， 不然无法反回去了

	for mybst != nil || mystack.Len() != 0 {
		for mybst != nil {
			mystack.PushBack(mybst) //首先左边压入栈中
			mybst = mybst.Left      //左边
		}
		v := mystack.Back() //取出节点
		top := v.Value.(*Node)
		// 循环到了尽头.
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && PreVisited == top.Left) || (PreVisited == top.Right) {
			res = append(res, top.Data) //压入数据
			PreVisited = top            //记录上一个节点
			mystack.Remove(v)           //处理完了在栈中删除
		} else {
			mybst = top.Right //右边循环
		}
	}
	return res
}

func (bst *BinaryTree) Levelshow() {
	bst.levelshow(bst.Root)
}

// 广度遍历
func (bst *BinaryTree) levelshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Front() //前面取出数据
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

func (bst *BinaryTree) Stackshow(n *Node) {
	myqueue := list.New() //新建一个list模拟队列
	myqueue.PushBack(n)   //后面压入数据
	for myqueue.Len() > 0 {
		left := myqueue.Back() //前面取出数据 ,此时是栈
		right := left.Value
		myqueue.Remove(left) //删除
		if v, ok := right.(*Node); ok && v != nil {
			fmt.Println(v.Data) //打印数据
			myqueue.PushBack(v.Left)
			myqueue.PushBack(v.Right)
		}
	}
}

// 最小公共祖先
func (bst *BinaryTree) FindlowerstAncestor(root *Node, nodea *Node, nodeb *Node) *Node {
	if root == nil {
		return nil
	}
	if root == nodea || root == nodeb {
		return root //有一个节点是根节点，
	}
	left := bst.FindlowerstAncestor(root.Left, nodea, nodeb)   //递归查找
	right := bst.FindlowerstAncestor(root.Right, nodea, nodeb) //递归查找
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}

// 递归二叉树深度
func (bst *BinaryTree) GetDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	lengthleft := bst.GetDepth(root.Left)
	rightlength := bst.GetDepth(root.Right)
	if lengthleft > rightlength {
		return lengthleft + 1
	} else {
		return rightlength + 1
	}
}
