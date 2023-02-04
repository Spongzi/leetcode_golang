package _11

/*
剑指 Offer II 011. 0 和 1 个数相同的子数组
地址: https://leetcode.cn/problems/A1NYOS/
基础思想,把0转化为-1然后把题目转化为和为0的最长数组长度
*/

func findMaxLength(nums []int) int {
	sumToIndex := map[int]int{}
	sumToIndex[0] = -1
	sum, maxLength := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if prevIndex, ok := sumToIndex[sum]; ok {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			sumToIndex[sum] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
