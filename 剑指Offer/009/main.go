package _09

/*
剑指 Offer II 009. 乘积小于 K 的子数组
地址: https://leetcode.cn/problems/ZVAVXX/
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	// 定义一个乘数, 一个左指针, 一个结果值
	product := 1
	left := 0
	ans := 0
	// 遍历右边的指针,想让右边的指针往前走
	for right := 0; right < len(nums); right++ {
		// 记录移动到当前数字的乘积值
		product *= nums[right]
		// 循环遍历直到出现符合题目情况的出现
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		// 得到的双指针区间内的数字, 怎么组合都是符合条件的(子数组),所以直接减去
		ans += right - left + 1
	}
	return ans
}
