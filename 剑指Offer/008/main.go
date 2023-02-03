package _08

import (
	"math"
)

/*
剑指 Offer II 008. 和大于等于 target 的最短子数组
地址: https://leetcode.cn/problems/2VG8Kg/?favorite=e8X3pBZi
*/

func minSubArrayLen(target int, nums []int) int {
	// 首先默认定义长度为int64的最大值
	ans := math.MaxInt64
	// 双指针进行判断
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum > target {
			ans = int(math.Min(float64(ans), float64(right-left+1)))
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
