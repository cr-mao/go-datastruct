package leetcode

import "fmt"

// import "fmt"

// func main() {
// 	fmt.Println(removeNum([]int{1, 2, 1, 3}, 1))
// }

func removeNum(num []int, val int) int {
	ret := 0
	length := len(num)
	for i := 0; i < length; i++ {
		if num[i] != val {
			num[ret] = num[i]
			ret++
		}
	}
	fmt.Println(num)
	return ret

}
