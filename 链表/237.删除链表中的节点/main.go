package 删除链表中的节点

/*
https://leetcode.cn/problems/delete-node-in-a-linked-list/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
