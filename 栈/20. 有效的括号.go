package 栈

import "strings"

/**
https://leetcode.cn/problems/valid-parentheses/
*/

// 暴力解法,效率太低...
func isValid(s string) bool {
	// 使用暴力的解法
	for strings.Contains(s, "{}") ||
		strings.Contains(s, "()") ||
		strings.Contains(s, "[]") {
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
	}
	return len(s) == 0
}

func isValid2(s string) bool {
	sLen := len(s)
	if sLen%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		'[': ']',
		'{': '}',
	}

	var stack []byte
	for i := 0; i < sLen; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}
