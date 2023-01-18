package 相交链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tep := headA; tep != nil; tep = tep.Next {
		vis[tep] = true
	}
	for temp := headB; temp != nil; temp = temp.Next {
		if vis[temp] {
			return temp
		}
	}
	return nil
}
