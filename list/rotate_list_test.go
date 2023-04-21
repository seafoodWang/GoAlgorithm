package list

import (
	"testing"
)

func TestRotateList(t *testing.T) {
	head := createListBySlice([]int{1})
	head = rotateRight(head, 1)
	PrintNodeList(head)
}

// 向右移动k个
func rotateRight(head *Node, k int) *Node {
	if head == nil || k == 0 {
		return head
	}

	length := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		length++
	}
	length = k % length
	slow := head
	fast := head
	for i := 0; i < length; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if slow == fast {
		return head
	}
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}
