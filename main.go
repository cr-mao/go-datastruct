package main

import (
    "datastruct/ArrayList"
    "datastruct/Myutils"
    "datastruct/Queue"
    "datastruct/StackArray"
    "datastruct/recurse"
    "fmt"
    "io/ioutil"
)
////arrayList
func ArrayListTest(){
    list := arrayList.NewArrayList()
    list.Append(2)
    list.Append(3)
    list.Append("222")
    fmt.Println(list)
    fmt.Println(list.TheSize)
    var p arrayList.List = arrayList.NewArrayList()
    for i := 0; i < 20; i++ {
        p.Append(i)
        fmt.Println(p)
    }
    p.Clear()
    fmt.Println(p)
}
////使用迭代器遍历数组
func IteratorTest() {
    list := arrayList.NewArrayList()
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
    mystack := arrayList.NewArrayListStack()
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

//数组栈迭代器
func ArrayListStackIteratorTest() {
    mystack := arrayList.NewArrayListStackX()
    mystack.Push(1)
    mystack.Push(2)
    mystack.Push(3)
    mystack.Push(4)
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    fmt.Println(mystack.Pop())
    mystack.Push(1)
    mystack.Push(2)
    mystack.Push(3)
    mystack.Push(4)
    for it := mystack.Myit; it.HasNext(); {
        val, _ := it.Next()
        fmt.Println(val)
    }
}

//递归文件夹
func RecursiveFolderTest() {
    path := "/Users/mac/code/golangproject/src/crblog.cc"
    files := []string{}                    //数组字符串
    files, _ = recurse.GetALL(path, files) //抓取所有文件
    for i := 0; i < len(files); i++ {      //打印路径
        fmt.Println(files[i])
    }
}

//栈模拟递归
func RecursiveFolderWithStackTest() {
    path := "/Users/mac/code/golangproject/src/crblog.cc"
    files := []string{} //数组字符串
    mystack := stackArray.NewStack()
    mystack.Push(path)
    for !mystack.IsEmpty() {
        path := mystack.Pop().(string)
        files = append(files, path)     //加入列表
        read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
        for _, fi := range read {
            if fi.IsDir() {
                fulldir := path + "/" + fi.Name() //构造新的路径
                mystack.Push(fulldir)

            } else {
                fulldir := path + "/" + fi.Name() //构造新的路径
                files = append(files, fulldir)    //追加路径
            }
        }

    }
    for i := 0; i < len(files); i++ { //打印
        fmt.Println(files[i])
    }
}

//递归文件夹并打印层级
func RecursiveFolderWithLevelTest() {
    path := "/Users/mac/code/golangproject/src/crblog.cc"
    recurse.GetALLX(path, 1)
}

//栈模拟递归文件并打印层级
func RecursiveFolderWithLevelStackTest() {
    path := "/Users/mac/code/golangproject/src/crblog.cc"
    files := []string{}                        //数组字符串
    files, _ = recurse.GetALLY(path, files, 1) //抓取所有文件
    for i := 0; i < len(files); i++ {          //打印路径
        fmt.Println(files[i])
    }
}

//队列测试
func QueueTest() {
    myq := queue.NewQueue()
    myq.EnQueue(1)
    myq.EnQueue(2)
    myq.EnQueue(3)
    myq.EnQueue(4)
    fmt.Println(myq.DeQueue())
    fmt.Println(myq.DeQueue())
    fmt.Println(myq.DeQueue())
    fmt.Println(myq.DeQueue())
    myq.EnQueue(14)
    myq.EnQueue(114)
    fmt.Println(myq.DeQueue())
    fmt.Println(myq.DeQueue())
    myq.EnQueue(11114)
    fmt.Println(myq.DeQueue())
}


func main(){

    //数组，迭代器，栈迭代器，递归文件夹，栈递归文件夹
    //ArrayListTest()
    //IteratorTest()
    //ArrayListStackTest()
    //ArrayListStackIteratorTest()
    //RecursiveFolderTest()
    //RecursiveFolderWithStackTest()
    //RecursiveFolderWithLevelTest()
    //RecursiveFolderWithLevelStackTest()
    //QueueTest()

    //选择排序
    //sortUtil.SelectSortTest()
    //冒泡排序
    //sortUtil.BubbleSortTest()
    //插入排序
    //sortUtil.InsertSortTest()
    //堆排序
    //sortUtil.HeapSortTest()
    //快速排序
    //  sortUtil.QuickSortTest()
    //奇偶排序
    //sortUtil.OddEvenTest()

   // searchUtils.BinarySearchTest()
   //searchUtils.CountFileLinesTest()
   // searchUtils.SequentialSearchTest()

   //快速排序 二分查找 用户数据检索
   //searchUtils.QuickBinarySearchTest()
    //自定义异常
    //myexception.MyexceptionTest()

    //searchUtils.SearchQQLine()
    //searchUtils.SearchQQ()
    //searchUtils.AdvanceQuickSortTest()
    //searchUtils.AdvanceQuickSortAndBinarySearchQQTest()
    //计数排序
    //searchUtils.CountNumSortTest()



}