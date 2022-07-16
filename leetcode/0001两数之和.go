package leetcode

/**
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        dic = {}
        for i, num in enumerate(nums):
            if target - num in dic:
                return [dic[target - num], i]
            dic[num] = i
*/
//https://blog.csdn.net/qq_32424059/article/details/89521067
// func main() {
// 	fmt.Println(twoSum([]int{2, 3, 5, 9}, 10))
// 	fmt.Println(twoSumMap([]int{2, 3, 5, 9}, 5))
// }

//a 给定数组
//c  targe
//返回 下标
func twoSum(a []int, c int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == c {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumMap(a []int, c int) []int {
	var m1 = make(map[int]int)
	for i := 0; i < len(a); i++ {
		if j, ok := m1[c-a[i]]; ok {
			return []int{j, i}
		}
		m1[a[i]] = i
	}
	return []int{}
}
