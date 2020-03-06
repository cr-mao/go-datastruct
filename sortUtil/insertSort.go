package sortUtil

import "fmt"

/*
1, 9 2 28 3 4
1 9 2 28 3 4
1 2 9 28 3 4
1 2 3 9 28  4
12 3 4 9 28

*/
//插入排序
func InsertForOne(arr []int) []int {
	//假定处理下标3的， 从下标2开始找，找到比 下标3的值大的位置插入进去，从小到大处理
	num := arr[3]
	j := 3 - 1
	for ; j >= 0 && arr[j] > num; j-- {
		arr[j+1] = arr[j]
	}
	arr[j+1] = num
	return arr
}

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > num; j-- { //找到小的位置插入进去，从小到大
			//for ; j >= 0 && arr[j] < num; j-- { //找到大的位置插入进去，从大到小
			arr[j+1] = arr[j]
		}
		arr[j+1] = num
	}
	return arr
}

func InsertSortTest() {
	var arr = []int{1, 2, 9, 8, 0, 3, 7, 4, 5, 6}
	fmt.Println(InsertForOne(arr)) //[1 2 8 9 0 3 7 4 5 6]
	fmt.Println(InsertSort(arr))
}
