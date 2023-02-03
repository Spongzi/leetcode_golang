package _06

/*
剑指 Offer II 006. 排序数组中两个数字之和
<a href="https://leetcode.cn/problems/kLl5u1/?favorite=e8X3pBZi"></a>
*/
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for i := 0; i < len(numbers); i++ {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{left, right}
}
