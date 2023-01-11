package 设计链表

type MyLinkedList struct {
	size int
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &Node{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	returnNode := this.head
	for i := 0; i <= index; i++ {
		returnNode = returnNode.Next
	}
	return returnNode.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	index = max(index, 0)
	this.size++
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &Node{Val: val, Next: pred.Next}
	pred.Next = toAdd
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	pred := this.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

type Node struct {
	Val  int
	Next *Node
}

func max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
