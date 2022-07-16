package leetcode

import "math"

// 整数翻转
// 1. 逆序输出， (暴力解法）
// 2. 顺序交互
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	if x == 0 {
		return 0
	}

	var result int
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		result = x%10 + result*10
		x = x / 10
	}
	return result
}
