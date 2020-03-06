package sortUtil

import "fmt"

//堆排序
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - i     //每次截取一段
		HeapSortMax(arr, lastmesslen) //每次截取0-lastmesslen-1元素，取到最大值放到arr第一个元素。    然后交换到  lastmesslen-1得位置
		fmt.Println(arr)
		arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
		fmt.Println("ex", arr)
	}
	return arr
}

//lenth=9 , depth=3
//8,7,6,5,4,3,1,2,0
//3,2,1,0
//3,7,8
//3 保存最大值
//2 最大值
//1 最大值
//0 最大值

func HeapSortMax(arr []int, length int) []int {
	//length:=len(arr)//数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		depth := length/2 - 1         // 深度， n  2*n+1,2*n+2 ,树的数量等于depth+1个
		for i := depth; i >= 0; i-- { //循环所有的三节点
			topmax := i //假定最大的在i的位置
			leftchild := 2*i + 1
			rightchild := 2*i + 2                                      //左右孩子的节点
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越界
				topmax = leftchild //如果左边比我大，记录最大
			}
			if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
				topmax = rightchild //如果右边比我大，记录最大
			}
			if topmax != i { //确保i的值就是最大
				arr[i], arr[topmax] = arr[topmax], arr[i]
			}
			//fmt.Println(i,arr)
		}
		return arr
	}
}
func HeapSortTest() {
	arr := []int{1, 15,9, 2, 8, 3, 7, 4, 6,11, 5, 10,20}
	//fmt.Println(HeapSortMax(arr,11))
	fmt.Println(HeapSort(arr))
}
