/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/7/4
**/
package leetcode

func ReverseStr(inputString string) string {
	data := []rune(inputString)
	length := len(data)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return string(data)
}
