package arrayList

type StackArrayX interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsFull() bool          //是否满了
	IsEmpty() bool         //是否空了
}

type StackX struct {
	myarray *ArrayList
	Myit    Iterator
}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()
	mystack.Myit = mystack.myarray.Iterator()
	return mystack
}
func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
}

func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *StackX) Pop() interface{} {
	if mystack.IsEmpty() {
		return nil
	}
	last := mystack.myarray.dataStore[mystack.myarray.TheSize-1]
	mystack.myarray.Delete(mystack.myarray.TheSize - 1)
	return last
}

func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}

func (mystack *StackX) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *StackX) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}
