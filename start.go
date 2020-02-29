package main

import (
    "datastruct/ArrayList"
    "fmt"
)

func ArrayListTest(){
    list :=ArrayList.NewArrayList()
    list.Append(2)
    list.Append(3)
    list.Append("222")
    fmt.Println(list)
    fmt.Println(list.TheSize)
    var p ArrayList.List = ArrayList.NewArrayList()
    for i := 0; i < 20; i++ {
        p.Append(i)
        fmt.Println(p)
    }
    p.Clear()
    fmt.Println(p)
}

func IteratorTest() {
    list := ArrayList.NewArrayList()
    list.Append(1)
    list.Append(2)
    list.Append(3)
    list.Append("222")
    it := list.Iterator()
    for it.HasNext() {
        val, _ := it.Next()
        fmt.Println(val)
    }
}


//栈结合ArrayList实现 (ArrayListStack)
func ArrayListStackTest() {
    mystack := ArrayList.NewArrayListStack()
    mystack.Push(1)
    mystack.Push(2)
    mystack.Push(3)
    mystack.Push(4)
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    mystack.Push(1)
    fmt.Println(mystack.Pop())
    mystack.Push(2)
    fmt.Println(mystack.Pop())
    mystack.Push(3)
    fmt.Println(mystack.Pop())
    mystack.Push(4)
    fmt.Println(mystack.Pop())
}

func main(){
    //ArrayListTest()
    //IteratorTest()
    ArrayListStackTest()

}