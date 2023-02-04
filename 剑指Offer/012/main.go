package _12

/*
剑指 Offer II 012. 左右两边子数组的和相等
地址: https://leetcode.cn/problems/tvdfij/solution/
*/

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		// leftSum-nums[i] == sum-leftSum
		if leftSum*2-nums[i] == sum {
			return i
		}
	}
	return -1
}
