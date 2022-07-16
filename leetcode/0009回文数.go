package leetcode 



// //回文数
// func main() {
// 	fmt.Println(isHuiwen(121))
// 	fmt.Println(isHuiwen(0))
// 	fmt.Println(isHuiwen(-121))
// }

func isHuiwen(a int) bool {
	if a < 0 || a != 0 && a%10 == 0 {
		return false
	}
	var res int
	b := a
	for a != 0 {
		res = res*10 + a%10
		a = a / 10
	}
	return res == b
}
