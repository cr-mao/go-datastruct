package arraylist

import (
	"fmt"
	"errors"
)

//定义接口
type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //获得第index元素
	Set(index int, newVal interface{}) error    //修改
	Insert(index int, newVal interface{}) error //插入
	Append(newVal interface{})                  //追加
	Clear()                                     //清空
	Delete(index int) error                     //删除元素
	String() string                             //返回字符串
}

//数据结构     可以存 整数，字符串，实数
type ArrayList struct {
	dataStore []interface{} //数组存储
	TheSize   int           //数组大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10) //开辟10个空间
	list.TheSize = 0
	return list
}

//返回数组的大小
func (list *ArrayList) Size() int {
	return list.TheSize //返回数组的大小
}

//获得元素
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}
func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newVal
	return nil
}

//追加
func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.TheSize++
}

//返回数组的字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

//清空
func (list *ArrayList) Clear() {
	//重新开辟内存空间
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

//删除
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //重新叠加 跳过index
	list.TheSize--
	return nil
}

//检查是否已经满了，满了就开辟2倍内存
func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		//必须先开辟内存， 第二个参数不能为0 不然复制有问题
		newDataStore := make([]interface{}, 2*cap(list.dataStore), 2*cap(list.dataStore))
		copy(newDataStore, list.dataStore)
		list.dataStore = newDataStore
	}
}

//插入
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	//判断内存是否已经满了
	list.checkIsFull()
	//最后的元素 向后移动一位啊。
	for allSize := list.TheSize; allSize > index; allSize-- {
		list.dataStore[allSize] = list.dataStore[allSize-1]
	}
	list.dataStore[index] = newVal
	list.TheSize++
	return nil

}
