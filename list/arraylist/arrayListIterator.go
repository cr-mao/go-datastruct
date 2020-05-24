package arraylist

import "errors"

type Iterator interface {
	HasNext() bool              //是否有下一个
	Next() (interface{}, error) //下一个
	Remove()                    //删除当前索引元素
	GetIndex() int              //得到当前索引
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}

type ArraylistIterator struct {
	list         *ArrayList //数组指针
	currentIndex int        //当前索引
}

func (it *ArraylistIterator) HasNext() bool { //是否有下一个
	return it.currentIndex < it.list.TheSize
}

//获得下一个元素
func (it *ArraylistIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}

//删除当前索引元素
func (it *ArraylistIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func (it *ArraylistIterator) GetIndex() int {
	return it.currentIndex
}

//构造迭代器
func (list *ArrayList) Iterator() Iterator {
	it := new(ArraylistIterator)
	it.list = list
	it.currentIndex = 0
	return it
}
