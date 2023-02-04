package _10

/*
剑指 Offer II 010. 和为 k 的子数组
地址: https://leetcode.cn/problems/QTMn0o/
*/

func subarraySum(nums []int, k int) int {
	sumToCount := map[int]int{}
	sumToCount[0] = 1
	sum, count := 0, 0
	for _, num := range nums {
		sum += num
		if _, ok := sumToCount[sum-k]; ok {
			count += sumToCount[sum-k]
		}
		sumToCount[sum] += 1
	}
	return count
}

func subarraySum1(nums []int, k int) int {
	count := 0
	for i, _ := range nums {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}
