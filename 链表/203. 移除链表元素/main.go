package 移除链表元素

/**
https://leetcode.cn/problems/remove-linked-list-elements/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	temp := dummyHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return dummyHead.Next
}
