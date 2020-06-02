package arraylist

type StackArrayIterator interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否空了
}


type StackIterator struct {
	myarray *ArrayList
	Myit    Iterator
}

func NewArrayListStackIterator() *StackIterator {
	mystack := new(StackIterator)
	mystack.myarray = NewArrayList()
	mystack.Myit = mystack.myarray.Iterator()
	return mystack
}
func (mystack *StackIterator) Clear() {
	mystack.myarray.Clear()
}

func (mystack *StackIterator) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *StackIterator) Pop() interface{} {
	if mystack.IsEmpty() {
		return nil
	}
	last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
	mystack.myarray.Delete(mystack.myarray.TheSize - 1)
	return last
}

func (mystack *StackIterator) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *StackIterator) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *StackIterator) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}
