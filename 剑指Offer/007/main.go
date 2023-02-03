package _07

import "sort"

/*
剑指 Offer II 007. 数组中和为 0 的三个数
地址: https://leetcode.cn/problems/1fGaJU/?favorite=e8X3pBZi
*/
func threeSum(nums []int) [][]int {
	// 首先定义一个结果数组
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)

	// 遍历数组
	for first := 0; first < n; first++ {
		// 需要遍历和上一次不一样的数字
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		// 定义新的target目标数
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			// 需要保证遍历到不同的数字
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 保证second 在 third的左边
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果出现了指针重合也没有发现三数之和为0的那么久不会再出现了
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
