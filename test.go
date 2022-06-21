package main

import "fmt"

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - i     //每次截取一段
		HeapSortMax(arr, lastmesslen) //arr[0] 就是最大的
		fmt.Println(arr)
		arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0] //把最大的放在最后面
		fmt.Println("ex", arr)
	}
	return arr
}

func HeapSortMax(arr []int, length int) []int {
	//length:=len(arr)//数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		depth := length/2 - 1         //深度(顶点的最大深度）， n  2*n+1,2*n+2 ,树的数量等于depth+1个
		for i := depth; i >= 0; i-- { //循环所有的三节点
			topmax := i //假定最大的在i的位置 ，
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
		}
		return arr
	}
}


func ReverseStr(inputString string) string {
	data :=[]rune(inputString)
	length :=len(data)
	for i,j:=0,length-1;i<j; i,j =i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
	return string(data)
}



func main() {
	reverseStrRes :=ReverseStr("abcd毛dkdk")
	fmt.Println(reverseStrRes)
}

