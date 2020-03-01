package searchUtils

import "fmt"


//二分查找法
func BinarySearch(data []int, search int) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (high + low) / 2
		midData := data[mid]
		if (search > midData) {
			low = mid + 1
		} else if (search < midData) {
			high = mid - 1
		} else {
			return mid
		}
	}
	// 没找到
	return -1
}

func BinarySearchTest() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 4))
}






