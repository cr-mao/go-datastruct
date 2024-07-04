package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var left int
	var right = 1
	var max = 1
	for right < len(s) {
		//滑动窗口
		temp := s[left:right]
		has := false
		for i := 0; i < len(temp); i++ {
			if temp[i] == s[right] {
				has = true
				continue
			}
		}
		if !has {
			right++
		} else {
			left++
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
	// 左指针，右指针， 当 右边指针元素 出现之前出现过的。 移动左指针
	// abbcdefa
	// 并且 之前出现过的位置 要大于 左指针
}
