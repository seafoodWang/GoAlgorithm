package list

import (
	"testing"
)

func TestDutchFlag(t *testing.T) {

	head := createListBySlice([]int{1, 5, 2, 9, 6, 4, 7})
	h := dutchFlag(head, 10)
	PrintNodeList(h)
}

func dutchFlag(node *Node, x int) *Node {
	var smallHead, smallTail, midHead, midTail, bigHead, bigTail *Node
	for node != nil {
		if node.Val < x {
			if smallHead == nil {
				smallHead = node
				smallTail = node
			} else {
				smallTail.Next = node
				smallTail = smallTail.Next
			}
		} else if node.Val > x {
			if bigHead == nil {
				bigHead = node
				bigTail = node
			} else {
				bigTail.Next = node
				bigTail = bigTail.Next
			}
		} else {
			if midHead == nil {
				midHead = node
				midTail = node
			} else {
				midTail.Next = node
				midTail = midTail.Next
			}
		}
		node = node.Next
	}
	//链接链表&返回头指针
	head := contact(smallHead, smallTail, midHead)
	head = contact(head, midTail, bigHead)
	return head
}

func contact(h1, t1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	} else {
		if h2 != nil {
			t1.Next = h2
		}
	}
	return h1
}
