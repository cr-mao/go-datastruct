package arraylist

type StackArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否空了
}

type Stack struct {
	myarray *ArrayList
	capsize int //栈最大空间（最大范围）
}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList()
	mystack.capsize = 10 //空间
	return mystack
}
func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
}

func (mystack *Stack) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *Stack) Pop() interface{} {
	if mystack.IsEmpty() {
		return nil
	}
	last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
	_ = mystack.myarray.Delete(mystack.myarray.TheSize - 1)
	return last
}

func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsFull() bool {
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	} else {
		return false
	}
}
