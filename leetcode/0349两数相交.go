package leetcode

/**
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
*/

func intersection(nums1 []int, nums2 []int) []int {
	var m1 = make(map[int]int)
	var m2 = make(map[int]int)
	var result = make([]int, 0)
	for _, v := range nums1 {
		m1[v] = 1
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			m2[v] = 1
		}
	}
	for v := range m2 {
		result = append(result, v)
	}
	return result

}
