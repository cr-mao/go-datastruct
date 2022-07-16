package leetcode

// func main() {

// }

func removeDuplicates(nums []int) int {
	idx := 0
	for idx < len(nums)-1 {
		if nums[idx] == nums[idx+1] {
			nums = append(nums[:idx], nums[idx+1:]...) // 若相等则pop掉当前值
		} else { // 否则move到下一位置继续做判断
			idx += 1
		}
	}
	return len(nums)
}
