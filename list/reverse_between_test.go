package list

import "testing"

func TestReverseBetweenList(t *testing.T) {
	head := createList(5)
	head = reverseBetween(head, 1, 5)
}

func reverseBetween(head *Node, left int, right int) *Node {
	if head == nil {
		return head
	}
	node := head
	var pre *Node
	index := 0
	for node != nil {
		if index+1 == left {
			pre = node
			pre.Next = nil
			node = node.Next
		} else {
			cur := node
			node = node.Next
			cur.Next = node
			if index == right {
				break
			}
		}
	}
	if pre != nil {
		pre.Next = node
	}

	return head
}
