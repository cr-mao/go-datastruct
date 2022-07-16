package leetcode

func dpFibCompress(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	one := 1
	two := 2
	var result int

	for i := 3; i <= n; i++ {
		// 省了维护数组的开销
		result = one + two
		one = two
		two = result
	}
	return result

}
